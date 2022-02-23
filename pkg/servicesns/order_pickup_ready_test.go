package servicesns_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/internal/pkg/test"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/logger"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicesns"
)

func TestPublishOrderPickupReady(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	err := servicesns.PublishOrderPickupReady(logger.NewEmptyLogger(), context.TODO(), "dev", &servicesns.EventOrderPickupReadyData{
		OrderID: "POS-810052",
		InTime:  "10 Minuten",
	})
	Expect(err).To(BeNil())
}
