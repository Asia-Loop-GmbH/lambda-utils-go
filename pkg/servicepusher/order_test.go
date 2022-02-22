package servicepusher_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/logger"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicepusher"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/test"
)

func TestPublishOrderCreated(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	err := servicepusher.PublishOrderCreated(logger.NewEmptyLogger(), "dev", &servicepusher.EventOrderCreatedData{})
	Expect(err).To(BeNil())
}

func TestPublishOrderDelivered(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	err := servicepusher.PublishOrderDelivered(logger.NewEmptyLogger(), "dev", &servicepusher.EventOrderDeliveredData{})
	Expect(err).To(BeNil())
}

func TestPublishOrderPOSPaymentStarted(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	err := servicepusher.PublishOrderPOSPaymentStarted(logger.NewEmptyLogger(), "dev", &servicepusher.EventOrderPOSPaymentStartedData{})
	Expect(err).To(BeNil())
}

func TestPublishOrderPOSPaymentPaid(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	err := servicepusher.PublishOrderPOSPaymentPaid(logger.NewEmptyLogger(), "dev", &servicepusher.EventOrderPOSPaymentPaidData{})
	Expect(err).To(BeNil())
}

func TestPublishGroupFinalized(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	err := servicepusher.PublishGroupFinalized(logger.NewEmptyLogger(), "dev", &servicepusher.EventGroupFinalizedData{})
	Expect(err).To(BeNil())
}

func TestPublishGroupDelivered(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	err := servicepusher.PublishGroupDelivered(logger.NewEmptyLogger(), "dev", &servicepusher.EventGroupDeliveredData{})
	Expect(err).To(BeNil())
}
