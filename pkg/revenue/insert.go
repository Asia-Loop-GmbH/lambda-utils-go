package revenue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/sirupsen/logrus"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/revenue/service"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicedynamodb"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicessm"
)

func refundExists(log *logrus.Entry, ctx context.Context, stage string, merchantRef string) (bool, error) {
	log.Infof("check revenue exists [%s]", merchantRef)

	dbClient, err := servicedynamodb.NewClient(log, ctx)
	if err != nil {
		return false, err
	}

	table, err := servicessm.GetParameter(log, ctx, stage, "/dynamo/revenue", false)
	if err != nil {
		return false, err
	}

	output, err := dbClient.Query(ctx, &dynamodb.QueryInput{
		TableName:              table,
		IndexName:              aws.String("PaymentId"),
		KeyConditionExpression: aws.String("PaymentId = :paymentId"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":paymentId": &types.AttributeValueMemberS{Value: merchantRef},
		},
	})
	if err != nil {
		return false, err
	}

	if len(output.Items) == 0 {
		return false, nil
	}

	for _, item := range output.Items {
		r := new(service.Revenue)
		err := attributevalue.UnmarshalMap(item, r)
		if err != nil {
			return false, err
		}
		if r.Type == service.RevenueTypeRefund {
			return true, nil
		}
	}

	return false, nil
}

func orderExists(log *logrus.Entry, ctx context.Context, stage string, merchantRef string) (bool, error) {
	log.Infof("check revenue exists [%s]", merchantRef)

	dbClient, err := servicedynamodb.NewClient(log, ctx)
	if err != nil {
		return false, err
	}

	table, err := servicessm.GetParameter(log, ctx, stage, "/dynamo/revenue", false)
	if err != nil {
		return false, err
	}

	output, err := dbClient.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: table,
		Key: map[string]types.AttributeValue{
			"Id": &types.AttributeValueMemberS{Value: merchantRef},
		},
	})
	if err != nil {
		return false, err
	}

	return output.Item != nil, nil
}

func insert(log *logrus.Entry, ctx context.Context, stage string, r *service.Revenue) error {
	log.Infof("insert revenue: %v", r)

	dbClient, err := servicedynamodb.NewClient(log, ctx)
	if err != nil {
		return err
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
		log.Infof("revenue [%s] exists, will be removed", r.ID)
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
	log.Infof("inserted revenue [%s]", r.ID)
	return nil
}
