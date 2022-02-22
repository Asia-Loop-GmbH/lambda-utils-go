package service_order_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/logger"
	serviceorder "github.com/asia-loop-gmbh/lambda-utils-go/v3/service-order"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/test"
)

func TestNextOrderInvoice(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))
	log := logger.NewEmptyLogger()
	ctx := context.Background()

	next, err := serviceorder.NextOrderInvoice(log, ctx, "dev")
	Expect(err).To(BeNil())
	log.Infof("%s", *next)
}
