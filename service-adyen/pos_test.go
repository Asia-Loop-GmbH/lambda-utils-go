package service_adyen_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/logger"
	serviceadyen "github.com/asia-loop-gmbh/lambda-utils-go/v3/service-adyen"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/test"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/text"
)

func TestNewTender(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	orderId := text.RandomString(6, true, true, true)
	Expect(serviceadyen.NewTender(logger.NewEmptyLogger(), "dev", "S1F2-000158213300585", orderId, 10.12)).To(BeNil())
}
