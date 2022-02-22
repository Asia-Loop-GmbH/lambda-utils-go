package mymongo_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/logger"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/mymongo"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/test"
)

func TestNextByStage(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	next, err := mymongo.Next(logger.NewEmptyLogger(), context.Background(), "dev", "test")
	Expect(err).To(BeNil())
	Expect(next > 0).To(BeTrue())
	mymongo.Disconnect(logger.NewEmptyLogger(), context.Background())
}
