package servicesns

import (
	"context"
	"fmt"

	"github.com/asia-loop-gmbh/lambda-utils-go/v4/pkg/servicessm"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	commoncontext "github.com/nam-truong-le/lambda-utils-go/pkg/context"
	"github.com/nam-truong-le/lambda-utils-go/pkg/logger"
)

type EventOrderPickupReadyData struct {
	OrderID string
	InTime  string
}

func PublishOrderPickupReady(ctx context.Context, data *EventOrderPickupReadyData) error {
	log := logger.FromContext(ctx)
	stage, ok := ctx.Value(commoncontext.FieldStage).(string)
	if !ok {
		return fmt.Errorf("undefined stage in context")
	}
	topic, err := servicessm.GetGlobalParameter(ctx, "/sns/order/pickup-ready/arn", false)
	if err != nil {
		log.Errorf("failed to get topic arn: %s", err)
		return err
	}
	c, err := getClient(ctx)
	if err != nil {
		return err
	}
	params := &sns.PublishInput{
		TopicArn:       topic,
		Message:        aws.String(fmt.Sprintf("order pickup ready [%s]: %s", data.OrderID, data.InTime)),
		MessageGroupId: aws.String(stage),
		MessageAttributes: map[string]types.MessageAttributeValue{
			"env": {
				DataType:    aws.String("String"),
				StringValue: aws.String(stage),
			},
			"orderId": {
				DataType:    aws.String("String"),
				StringValue: aws.String(data.OrderID),
			},
			"inTime": {
				DataType:    aws.String("String"),
				StringValue: aws.String(data.InTime),
			},
		},
	}
	if _, err := c.Publish(ctx, params); err != nil {
		log.Errorf("failed to publish")
		return err
	}
	return nil
}
