package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/sirupsen/logrus"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicedynamodb"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicessm"
)

func QueryByPaymentID(log *logrus.Entry, ctx context.Context, stage, paymentID string) ([]Revenue, error) {
	client, err := servicedynamodb.NewClient(log, ctx)
	if err != nil {
		return nil, err
	}
	table, err := servicessm.GetParameter(log, ctx, stage, "/dynamo/revenue", false)
	if err != nil {
		return nil, err
	}
	out, err := client.Query(ctx, &dynamodb.QueryInput{
		TableName:              table,
		IndexName:              aws.String("PaymentId"),
		KeyConditionExpression: aws.String("PaymentId = :paymentId"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":paymentId": &types.AttributeValueMemberS{Value: paymentID},
		},
	})
	rs := make([]Revenue, 0)
	if err := attributevalue.UnmarshalListOfMaps(out.Items, &rs); err != nil {
		return nil, err
	}
	return rs, nil
}

func GetByID(log *logrus.Entry, ctx context.Context, stage, orderID string) (*Revenue, error) {
	dbClient, err := servicedynamodb.NewClient(log, ctx)
	if err != nil {
		return nil, err
	}

	table, err := servicessm.GetParameter(log, ctx, stage, "/dynamo/revenue", false)
	if err != nil {
		return nil, err
	}

	out, err := dbClient.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: table,
		Key: map[string]types.AttributeValue{
			"Id": &types.AttributeValueMemberS{Value: orderID},
		},
	})
	if err != nil {
		return nil, err
	}

	if out.Item == nil {
		return nil, ErrorNotFound
	}

	r := new(Revenue)
	if err := attributevalue.UnmarshalMap(out.Item, r); err != nil {
		return nil, err
	}
	return r, nil
}
