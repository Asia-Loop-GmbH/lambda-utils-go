package revenue

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/internal/pkg/test"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/logger"
)

func TestOrderExistTrue(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	r, err := orderExists(logger.NewEmptyLogger(), context.TODO(), "dev", "263")
	Expect(err).To(BeNil())
	Expect(r).To(BeTrue())
}

func TestOrderExistFalse(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	r, err := orderExists(logger.NewEmptyLogger(), context.TODO(), "dev", "999")
	Expect(err).To(BeNil())
	Expect(r).To(BeFalse())
}

func TestRefundExistTrue(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	r1, err := refundExists(logger.NewEmptyLogger(), context.TODO(), "dev", "263")
	Expect(err).To(BeNil())
	Expect(r1).To(BeTrue())
}

func TestRefundExistFalse(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	r1, err := refundExists(logger.NewEmptyLogger(), context.TODO(), "dev", "259")
	Expect(err).To(BeNil())
	Expect(r1).To(BeFalse())
}
