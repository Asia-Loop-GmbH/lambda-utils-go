package servicesns_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/internal/pkg/test"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/logger"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicesns"
)

func TestPublishOrderFinalized(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	id, _ := primitive.ObjectIDFromHex("622a5275b73e4d6262fd8acf")
	err := servicesns.PublishOrderFinalized(logger.NewEmptyLogger(), context.TODO(), "dev", &servicesns.EventOrderFinalizedData{
		ID: id,
	})
	Expect(err).To(BeNil())
}
