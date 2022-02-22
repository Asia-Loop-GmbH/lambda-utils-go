package service_adyen_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/logger"
	serviceadyen "github.com/asia-loop-gmbh/lambda-utils-go/v3/service-adyen"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/test"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/text"
)

func TestNewDropInPayment_Success(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))
	response, err := serviceadyen.NewDropInPayment(
		logger.NewEmptyLogger(),
		"dev",
		"10.23",
		text.RandomString(10, true, true, true),
		"https://admin2-dev.asia-loop.com",
	)
	Expect(err).To(BeNil())
	Expect(response).NotTo(BeNil())
}
