package service_adyen_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/asia-loop-gmbh/lambda-utils-go/logger"
	serviceadyen "github.com/asia-loop-gmbh/lambda-utils-go/service-adyen"
	"github.com/asia-loop-gmbh/lambda-utils-go/test"
	"github.com/asia-loop-gmbh/lambda-utils-go/text"
)

func TestNewTender(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	orderId := text.RandomString(6, true, true, true)
	Expect(serviceadyen.NewTender(logger.NewEmptyLogger(), "dev", "S1F2-000158213300585", orderId, 10.12)).To(BeNil())
}
