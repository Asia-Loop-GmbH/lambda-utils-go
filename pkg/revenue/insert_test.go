package revenue

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/internal/pkg/test"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/logger"
)

func TestOrderExists(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	r1, err := orderExists(logger.NewEmptyLogger(), context.TODO(), "dev", "12")
	Expect(err).To(BeNil())
	Expect(r1).To(BeFalse())
}
