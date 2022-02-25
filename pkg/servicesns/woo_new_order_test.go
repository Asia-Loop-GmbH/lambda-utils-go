package servicesns_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/internal/pkg/test"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/logger"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicesns"
)

func TestPublishWooNewOrder(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	err := servicesns.PublishWooNewOrder(logger.NewEmptyLogger(), context.TODO(), "dev", &servicesns.EventWooNewOrderData{
		ID: 1234,
	})
	Expect(err).To(BeNil())
}
