package order_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/internal/pkg/test"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/logger"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicewoo/order"
)

func TestGet(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	o, err := order.Get(logger.NewEmptyLogger(), context.TODO(), "dev", 123)
	Expect(err).To(BeNil())
	Expect(o.ID).To(Equal(123))
}
