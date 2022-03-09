package revenue_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/internal/pkg/test"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/logger"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/revenue"
)

func TestSyncCorporateOrder(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	err := revenue.SyncCorporateOrder(logger.NewEmptyLogger(), context.TODO(), "dev", "POS-9C563XKE")
	Expect(err).To(BeNil())
}
