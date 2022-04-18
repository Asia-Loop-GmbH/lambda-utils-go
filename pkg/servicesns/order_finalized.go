package servicesns

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicessm"
)

type EventOrderFinalizedData struct {
	ID primitive.ObjectID
}

func PublishOrderFinalized(log *logrus.Entry, ctx context.Context, stage string, data *EventOrderFinalizedData) error {
	topic, err := servicessm.GetParameter(log, ctx, "all", "/sns/order/finalized/arn", false)
	if err != nil {
		log.Errorf("failed to get topic arn: %s", err)
		return err
	}
	client, err := getClient(log, ctx)
	if err != nil {
		return err
	}
	params := &sns.PublishInput{
		TopicArn:       topic,
		Message:        aws.String(fmt.Sprintf("order finalized [%s]", data.ID)),
		MessageGroupId: aws.String(stage),
		MessageAttributes: map[string]types.MessageAttributeValue{
			"env": {
				DataType:    aws.String("String"),
				StringValue: aws.String(stage),
			},
			"orderId": {
				DataType:    aws.String("String"),
				StringValue: aws.String(data.ID.Hex()),
			},
		},
	}
	if _, err := client.Publish(ctx, params); err != nil {
		log.Errorf("failed to publish")
		return err
	}
	return nil
}
