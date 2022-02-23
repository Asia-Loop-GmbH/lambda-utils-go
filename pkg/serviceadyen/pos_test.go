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

func TestNewTender(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	orderId := random.String(6, true, true, true)
	Expect(serviceadyen.NewTender(logger.NewEmptyLogger(), context.TODO(), "dev", "S1F2-000158213300585", orderId, 10.12)).To(BeNil())
}
