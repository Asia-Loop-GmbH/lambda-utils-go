package revenue

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/sirupsen/logrus"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicedynamodb"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicessm"
)

const (
	nonWooRefundSuffix = "--non-woo-refund"
)

func fakeNonWooRefund(log *logrus.Entry, ctx context.Context, stage, id string) error {
	client, err := servicedynamodb.NewClient(log, ctx)
	if err != nil {
		return err
	}
	table, err := servicessm.GetParameter(log, ctx, stage, "/dynamo/revenue", false)
	if err != nil {
		return err
	}

	out, err := client.Query(ctx, &dynamodb.QueryInput{
		TableName:              table,
		IndexName:              aws.String("PaymentId"),
		KeyConditionExpression: aws.String("PaymentId = :paymentId"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":paymentId": &types.AttributeValueMemberS{Value: id},
		},
	})

	if len(out.Items) != 1 {
		return fmt.Errorf("expected only one revenue item from [%s], got: %v", id, out.Items)
	}

	r := new(Revenue)
	err = attributevalue.UnmarshalMap(out.Items[0], r)
	if err != nil {
		return err
	}

	if r.Type != RevenueTypeOrder {
		return fmt.Errorf("expected only order from ref [%s], got: %v", id, r)
	}

	refund := Revenue{
		ID:             r.ID + nonWooRefundSuffix,
		PaymentID:      r.PaymentID,
		CreatedAt:      r.CreatedAt, // TODO: get date form adyen
		Type:           RevenueTypeRefund,
		ShippingMethod: r.ShippingMethod,
		Store:          r.Store,
		Source:         r.Source,
		Company:        r.Company,
		Net7:           "-" + r.Net7,
		Tax7:           "-" + r.Tax7,
		Tip:            "-" + r.Tip,
	}

	return insert(log, ctx, stage, &refund)
}
