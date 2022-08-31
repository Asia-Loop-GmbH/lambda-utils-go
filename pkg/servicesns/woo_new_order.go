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

type EventWooNewOrderData struct {
	ID int // woocommerce id
}

func PublishWooNewOrder(ctx context.Context, data *EventWooNewOrderData) error {
	log := logger.FromContext(ctx)
	stage, ok := ctx.Value(commoncontext.FieldStage).(string)
	if !ok {
		return fmt.Errorf("undefined stage in context")
	}
	topic, err := servicessm.GetGlobalParameter(ctx, "/sns/woo/order/new", false)
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
		Message:        aws.String(fmt.Sprintf("new woo order [%d]", data.ID)),
		MessageGroupId: aws.String(stage),
		MessageAttributes: map[string]types.MessageAttributeValue{
			"env": {
				DataType:    aws.String("String"),
				StringValue: aws.String(stage),
			},
			"id": {
				DataType:    aws.String("String"),
				StringValue: aws.String(fmt.Sprintf("%d", data.ID)),
			},
		},
	}
	if _, err := c.Publish(ctx, params); err != nil {
		log.Errorf("failed to publish")
		return err
	}
	return nil
}
