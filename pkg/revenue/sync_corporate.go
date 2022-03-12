package revenue

import (
	"context"
	"fmt"

	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/dbadmin"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/revenue/service"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicemongo"
)

const (
	corporateSuffix = "--corporate"
)

func SyncCorporateOrderByUUID(log *logrus.Entry, ctx context.Context, stage, id string) error {
	defer servicemongo.Disconnect(log, ctx)

	colOrder, err := servicemongo.AdminCollection(log, ctx, stage, dbadmin.CollectionOrder)
	if err != nil {
		return err
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	findOrder := colOrder.FindOne(ctx, bson.M{"_id": objectID})
	o := new(dbadmin.Order)
	if err := findOrder.Decode(o); err != nil {
		return err
	}

	return syncCorporateOrder(log, ctx, stage, o)
}

func SyncCorporateOrder(log *logrus.Entry, ctx context.Context, stage, orderID string) error {
	defer servicemongo.Disconnect(log, ctx)

	colOrder, err := servicemongo.AdminCollection(log, ctx, stage, dbadmin.CollectionOrder)
	if err != nil {
		return err
	}

	findOrder := colOrder.FindOne(ctx, bson.M{"orderId": orderID})
	o := new(dbadmin.Order)
	if err := findOrder.Decode(o); err != nil {
		return err
	}

	return syncCorporateOrder(log, ctx, stage, o)
}

func syncCorporateOrder(log *logrus.Entry, ctx context.Context, stage string, o *dbadmin.Order) error {
	colStore, err := servicemongo.AdminCollection(log, ctx, stage, dbadmin.CollectionStore)
	if err != nil {
		return err
	}

	findStore := colStore.FindOne(ctx, bson.M{"_id": o.Store})
	s := new(dbadmin.Store)
	if err := findStore.Decode(s); err != nil {
		return err
	}

	net7 := decimal.Zero
	tax7 := decimal.Zero

	for _, item := range o.Items {
		switch item.TaxClass {
		case service.WooTaxClass7:
			net, err := decimal.NewFromString(item.Net)
			if err != nil {
				return err
			}
			tax, err := decimal.NewFromString(item.Tax)
			if err != nil {
				return err
			}
			net7 = net7.Add(net)
			tax7 = tax7.Add(tax)
		default:
			return fmt.Errorf("tax class [%s] not registered", item.TaxClass)
		}
	}

	net, err := decimal.NewFromString(o.Net)
	if err != nil {
		return err
	}
	tax, err := decimal.NewFromString(o.Tax)
	if err != nil {
		return err
	}

	if !net7.Equal(net) {
		return fmt.Errorf("expected net [%s] from order [%s], calculated [%s]", net, o.OrderID, net7)
	}
	if !tax7.Equal(tax) {
		return fmt.Errorf("expected tax [%s] from order [%s], calculated [%s]", tax, o.OrderID, tax7)
	}

	total, err := decimal.NewFromString(o.Total)
	if err != nil {
		return err
	}
	tip := decimal.Zero
	totalCheck := net7.Add(tax7).Add(tip)
	if !totalCheck.Equal(total) {
		return fmt.Errorf("expected total [%s] from order [%s], calculated [%s]", total, o.OrderID, totalCheck)
	}

	// split
	// we need more complex computation if we have different taxes
	// currently we have only 7%, it makes things easier
	employeeTotal, err := decimal.NewFromString(o.PaidTotal)
	if err != nil {
		return err
	}
	employeeNet := employeeTotal.DivRound(decimal.NewFromFloat(1.07), 2)
	employeeTax := employeeTotal.Sub(employeeNet)

	companyNet := net7.Sub(employeeNet)
	companyTax := tax7.Sub(employeeTax)

	createdAtString, err := service.TimeToDynamoString(o.CreatedAt)
	if err != nil {
		return err
	}

	employeeRevenue := service.Revenue{
		ID:             o.OrderID,
		PaymentID:      o.ID.Hex(),
		CreatedAt:      *createdAtString,
		ShippingMethod: o.ShippingMethod,
		Store:          s.Configuration.WPStoreKey,
		Source:         service.RevenueSourceCorporate,
		Company:        o.CompanyKey,
		Type:           service.RevenueTypeOrder,
		Net7:           employeeNet.StringFixed(2),
		Tax7:           employeeTax.StringFixed(2),
		Tip:            tip.StringFixed(2),
	}

	if err := insert(log, ctx, stage, &employeeRevenue); err != nil {
		return err
	}

	companyRevenue := service.Revenue{
		ID:             o.OrderID + corporateSuffix,
		PaymentID:      o.ID.Hex() + corporateSuffix,
		CreatedAt:      *createdAtString,
		ShippingMethod: o.ShippingMethod,
		Store:          s.Configuration.WPStoreKey,
		Source:         service.RevenueSourceCorporate,
		Company:        o.CompanyKey,
		Type:           service.RevenueTypeOrder,
		Net7:           companyNet.StringFixed(2),
		Tax7:           companyTax.StringFixed(2),
		Tip:            tip.StringFixed(2),
	}

	if err := insert(log, ctx, stage, &companyRevenue); err != nil {
		return err
	}

	return nil
}
