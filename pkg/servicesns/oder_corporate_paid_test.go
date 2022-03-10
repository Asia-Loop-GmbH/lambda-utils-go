package servicesns_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/internal/pkg/test"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/logger"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicesns"
)

func TestPublishOrderCorporatePaid(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	err := servicesns.PublishOrderCorporatePaid(logger.NewEmptyLogger(), context.TODO(), "dev", &servicesns.EventOrderCorporatePaidData{
		OrderID: "POS-9C563XKE",
	})
	Expect(err).To(BeNil())
}
