package revenue_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/internal/pkg/test"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/logger"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/revenue"
)

func TestSyncPOSOrder(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	err := revenue.SyncPOSOrder(logger.NewEmptyLogger(), context.TODO(), "dev", "POS-2ZFJ9973")
	Expect(err).To(BeNil())
}
