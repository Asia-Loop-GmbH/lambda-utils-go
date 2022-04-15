package revenue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/sirupsen/logrus"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/revenue/service"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicedynamodb"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicessm"
)

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
