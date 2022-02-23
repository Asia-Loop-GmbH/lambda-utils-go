package servicemongo_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/internal/pkg/test"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/logger"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicemongo"
)

func TestNextByStage(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	next, err := servicemongo.Next(logger.NewEmptyLogger(), context.Background(), "dev", "test")
	Expect(err).To(BeNil())
	Expect(next > 0).To(BeTrue())
	servicemongo.Disconnect(logger.NewEmptyLogger(), context.Background())
}
