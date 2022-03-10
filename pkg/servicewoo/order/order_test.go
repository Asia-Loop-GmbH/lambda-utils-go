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

func TestGetRefunds(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	refunds, err := order.GetRefunds(logger.NewEmptyLogger(), context.TODO(), "dev", 245)
	Expect(err).To(BeNil())
	Expect(len(refunds)).To(Equal(1))
	Expect(refunds[0].ID).To(Equal(262))
}
