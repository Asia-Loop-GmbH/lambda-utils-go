package slots_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/internal/pkg/test"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/logger"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicewoo/slots"
)

func TestGetSlots_Success(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	s, err := slots.GetSlots(logger.NewEmptyLogger(), context.TODO(), "dev")

	Expect(err).To(BeNil())
	Expect(s).NotTo(BeNil())
}
