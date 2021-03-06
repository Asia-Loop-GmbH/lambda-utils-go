package servicepusher

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/internal/pkg/test"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/logger"
)

func TestGetClient(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))
	log := logger.NewEmptyLogger()
	ctx := context.TODO()

	client, err := getClient(log, ctx, "dev")
	Expect(err).To(BeNil())
	Expect(client).To(Not(BeNil()))

	client, err = getClient(log, ctx, "dev")
	Expect(err).To(BeNil())
	Expect(client).To(Not(BeNil()))
}
