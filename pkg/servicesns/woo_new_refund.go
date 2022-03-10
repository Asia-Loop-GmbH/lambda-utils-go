package servicesns

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/sirupsen/logrus"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicessm"
)

type EventWooNewRefundData struct {
	ID int // woocommerce id
}

func PublishWooNewRefund(log *logrus.Entry, ctx context.Context, stage string, data *EventWooNewRefundData) error {
	topic, err := servicessm.GetParameter(log, ctx, "all", "/sns/woo/refund/new", false)
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
		Message:        aws.String(fmt.Sprintf("new woo refund for order [%d]", data.ID)),
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
	if _, err := client.Publish(ctx, params); err != nil {
		log.Errorf("failed to publish")
		return err
	}
	return nil
}
