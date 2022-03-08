package revenue

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/dbadmin"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicedynamodb"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicessm"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicewoo/order"
)

func SyncWooOrder(log *logrus.Entry, ctx context.Context, stage string, id int) error {
	o, err := order.Get(log, ctx, stage, id)
	if err != nil {
		return err
	}
	dbClient, err := servicedynamodb.NewClient(log, ctx)
	if err != nil {
		return err
	}

	net7 := decimal.Zero
	tax7 := decimal.Zero
	for _, item := range o.LineItems {
		switch item.TaxClass {
		case WooTaxClass7:
			net, err := decimal.NewFromString(item.Total)
			if err != nil {
				return err
			}
			tax, err := decimal.NewFromString(item.TotalTax)
			if err != nil {
				return err
			}
			net7 = net7.Add(net)
			tax7 = tax7.Add(tax)
		default:
			return fmt.Errorf("tax class [%s] not registered", item.TaxClass)
		}
	}

	tip := decimal.Zero
	for _, item := range o.FeeLines {
		switch item.Name {
		case WooFeeTip:
			tipTax, err := decimal.NewFromString(item.TotalTax)
			if err != nil {
				return err
			}
			if !tipTax.IsZero() {
				return fmt.Errorf("tip tax from order [%d] must be zero, got [%s]", o.ID, tipTax.StringFixed(2))
			}
			tipTotal, err := decimal.NewFromString(item.Total)
			if err != nil {
				return err
			}
			tip = tip.Add(tipTotal)
		default:
			return fmt.Errorf("fee name [%s] not registered", item.Name)
		}
	}

	total, err := decimal.NewFromString(o.Total)
	if err != nil {
		return err
	}
	checkTotal := net7.Add(tax7).Add(tip)
	if !total.Equal(checkTotal) {
		return fmt.Errorf("expected total [%s] from oder [%d], calculated [%s]", total.StringFixed(2), o.ID, checkTotal.StringFixed(2))
	}

	totalTax, err := decimal.NewFromString(o.TotalTax)
	if err != nil {
		return err
	}
	checkTotalTax := tax7
	if !totalTax.Equal(checkTotalTax) {
		return fmt.Errorf("expected total tax [%s] from oder [%d], calculated [%s]", totalTax.StringFixed(2), o.ID, checkTotalTax.StringFixed(2))
	}

	created, err := o.GetDateCreated()
	if err != nil {
		return err
	}

	r := Revenue{
		ID:             fmt.Sprintf("%d", o.ID),
		PaymentID:      fmt.Sprintf("%d", o.ID),
		CreatedAt:      created,
		ShippingMethod: dbadmin.OrderShippingMethod(o.ShippingLines[0].MethodID),
		Store:          o.GetStoreKey(),
		Source:         RevenueSourceOnline,
		Company:        "",
		Type:           RevenueTypeOrder,
		Net7:           net7.StringFixed(2),
		Tax7:           tax7.StringFixed(2),
		Tip:            tip.StringFixed(2),
	}
	item, err := attributevalue.MarshalMap(r)
	if err != nil {
		return err
	}

	table, err := servicessm.GetParameter(log, ctx, stage, "/dynamo/revenue", false)
	if err != nil {
		return err
	}

	output, err := dbClient.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: table,
		Key: map[string]types.AttributeValue{
			"Id": &types.AttributeValueMemberS{Value: r.ID},
		},
	})
	if err != nil {
		return err
	}
	if output.Item != nil {
		log.Infof("order [%s] exists, will be removed", r.ID)
		_, err := dbClient.DeleteItem(ctx, &dynamodb.DeleteItemInput{
			TableName: table,
			Key: map[string]types.AttributeValue{
				"Id": &types.AttributeValueMemberS{Value: r.ID},
			},
		})
		if err != nil {
			return err
		}
	}

	_, err = dbClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: table,
		Item:      item,
	})
	if err != nil {
		return err
	}
	return nil
}