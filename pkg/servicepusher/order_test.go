package servicepusher_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/internal/pkg/test"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/logger"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicepusher"
)

func TestPublishOrderCreated(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	err := servicepusher.PublishOrderCreated(logger.NewEmptyLogger(), context.TODO(), "dev", &servicepusher.EventOrderCreatedData{})
	Expect(err).To(BeNil())
}

func TestPublishOrderDelivered(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	err := servicepusher.PublishOrderDelivered(logger.NewEmptyLogger(), context.TODO(), "dev", &servicepusher.EventOrderDeliveredData{})
	Expect(err).To(BeNil())
}

func TestPublishOrderPOSPaymentStarted(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	err := servicepusher.PublishOrderPOSPaymentStarted(logger.NewEmptyLogger(), context.TODO(), "dev", &servicepusher.EventOrderPOSPaymentStartedData{})
	Expect(err).To(BeNil())
}

func TestPublishOrderPOSPaymentPaid(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	err := servicepusher.PublishOrderPOSPaymentPaid(logger.NewEmptyLogger(), context.TODO(), "dev", &servicepusher.EventOrderPOSPaymentPaidData{})
	Expect(err).To(BeNil())
}

func TestPublishGroupFinalized(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	err := servicepusher.PublishGroupFinalized(logger.NewEmptyLogger(), context.TODO(), "dev", &servicepusher.EventGroupFinalizedData{})
	Expect(err).To(BeNil())
}

func TestPublishGroupDelivered(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	err := servicepusher.PublishGroupDelivered(logger.NewEmptyLogger(), context.TODO(), "dev", &servicepusher.EventGroupDeliveredData{})
	Expect(err).To(BeNil())
}
