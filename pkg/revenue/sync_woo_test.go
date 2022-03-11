package revenue_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/internal/pkg/test"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/logger"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/revenue"
)

func TestSyncWooOrder(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	err := revenue.SyncWooOrder(logger.NewEmptyLogger(), context.TODO(), "dev", 259)
	Expect(err).To(BeNil())
}

//func TestSyncWooOrderPROD(t *testing.T) {
//	RegisterFailHandler(test.FailedHandler(t))
//
//	err := revenue.SyncWooOrder(logger.NewEmptyLogger(), context.TODO(), "prod", 32121)
//	Expect(err).To(BeNil())
//}
