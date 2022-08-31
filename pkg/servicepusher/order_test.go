package servicepusher_test

import (
	"context"
	"testing"

	commoncontext "github.com/nam-truong-le/lambda-utils-go/pkg/context"
	"github.com/stretchr/testify/assert"

	"github.com/asia-loop-gmbh/lambda-utils-go/v4/pkg/servicepusher"
)

func TestPublishOrderCreated(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	ctx := context.WithValue(context.TODO(), commoncontext.FieldStage, "dev")
	err := servicepusher.PublishOrderCreated(ctx, &servicepusher.EventOrderCreatedData{})
	assert.NoError(t, err)
}

func TestPublishOrderDelivered(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	ctx := context.WithValue(context.TODO(), commoncontext.FieldStage, "dev")
	err := servicepusher.PublishOrderDelivered(ctx, &servicepusher.EventOrderDeliveredData{})
	assert.NoError(t, err)
}

func TestPublishOrderPOSPaymentStarted(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	ctx := context.WithValue(context.TODO(), commoncontext.FieldStage, "dev")
	err := servicepusher.PublishOrderPOSPaymentStarted(ctx, &servicepusher.EventOrderPOSPaymentStartedData{})
	assert.NoError(t, err)
}

func TestPublishOrderPOSPaymentPaid(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	ctx := context.WithValue(context.TODO(), commoncontext.FieldStage, "dev")
	err := servicepusher.PublishOrderPOSPaymentPaid(ctx, &servicepusher.EventOrderPOSPaymentPaidData{})
	assert.NoError(t, err)
}

func TestPublishGroupFinalized(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	ctx := context.WithValue(context.TODO(), commoncontext.FieldStage, "dev")
	err := servicepusher.PublishGroupFinalized(ctx, &servicepusher.EventGroupFinalizedData{})
	assert.NoError(t, err)
}

func TestPublishGroupDelivered(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	ctx := context.WithValue(context.TODO(), commoncontext.FieldStage, "dev")
	err := servicepusher.PublishGroupDelivered(ctx, &servicepusher.EventGroupDeliveredData{})
	assert.NoError(t, err)
}
