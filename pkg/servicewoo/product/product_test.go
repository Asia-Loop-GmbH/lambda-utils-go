package product_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/internal/pkg/test"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/logger"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicewoo/product"
)

func TestGet(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	products, err := product.Get(logger.NewEmptyLogger(), context.TODO(), "dev")
	Expect(err).To(BeNil())
	Expect(len(products) > 0).To(BeTrue())
}
