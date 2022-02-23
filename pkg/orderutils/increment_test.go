package orderutils_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/internal/pkg/test"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/logger"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/orderutils"
)

func TestNextOrderInvoice(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))
	log := logger.NewEmptyLogger()
	ctx := context.Background()

	next, err := orderutils.NextOrderInvoice(log, ctx, "dev")
	Expect(err).To(BeNil())
	log.Infof("%s", *next)
}
