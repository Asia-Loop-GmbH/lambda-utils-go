package serviceadyen_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/internal/pkg/test"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/logger"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/random"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/serviceadyen"
)

func TestNewDropInPayment_Success(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))
	response, err := serviceadyen.NewDropInPayment(
		logger.NewEmptyLogger(),
		context.TODO(),
		"dev",
		"10.23",
		random.String(10, true, true, true),
		"https://admin2-dev.asia-loop.com",
	)
	Expect(err).To(BeNil())
	Expect(response).NotTo(BeNil())
}
