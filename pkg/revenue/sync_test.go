package revenue_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/internal/pkg/test"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/logger"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/revenue"
)

func TestCheckOrder(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	err := revenue.CheckOrder(logger.NewEmptyLogger(), context.TODO(), "dev", "259")
	Expect(err).To(BeNil())
	err = revenue.CheckOrder(logger.NewEmptyLogger(), context.TODO(), "dev", "62235c6c8ed20c2774179f40")
	Expect(err).To(BeNil())
	err = revenue.CheckOrder(logger.NewEmptyLogger(), context.TODO(), "dev", "62229b6c957a5b919f8360fb")
	Expect(err).To(BeNil())
}
