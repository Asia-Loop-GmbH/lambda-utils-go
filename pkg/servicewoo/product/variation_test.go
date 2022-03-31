package product_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/internal/pkg/test"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/logger"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicewoo/product"
)

func TestGetVariation(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	variations, err := product.GetVariation(logger.NewEmptyLogger(), context.TODO(), "dev", 24)
	Expect(err).To(BeNil())
	Expect(len(variations) > 0).To(BeTrue())
}
