package servicepusher

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/logger"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/test"
)

func TestGetClient(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	client, err := getClient(logger.NewEmptyLogger(), "dev")
	Expect(err).To(BeNil())
	Expect(client).To(Not(BeNil()))

	client, err = getClient(logger.NewEmptyLogger(), "dev")
	Expect(err).To(BeNil())
	Expect(client).To(Not(BeNil()))
}
