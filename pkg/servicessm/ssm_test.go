package servicessm_test

import (
	"context"
	"os"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/internal/pkg/test"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/logger"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicessm"
)

func TestGetParameter_FromEnv(t *testing.T) {
	dummyValue := "foo"
	RegisterFailHandler(test.FailedHandler(t))
	err := os.Setenv("AL_SSM_ALL_MONGO_HOST", dummyValue)
	Expect(err).To(BeNil())
	v, err := servicessm.GetParameter(logger.NewEmptyLogger(), context.TODO(), "all", "/mongo/host", false)

	Expect(*v).To(Equal(dummyValue))
}

func TestGetParameter_EnvNotExists(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))
	err := os.Unsetenv("AL_SSM_ALL_MONGO_HOST")
	Expect(err).To(BeNil())
	v, err := servicessm.GetParameter(logger.NewEmptyLogger(), context.TODO(), "all", "/mongo/host", false)

	Expect(*v).To(Equal("asia-loop-admin.yncuk.mongodb.net"))
}

func TestGetParameter_EnvEmptyString(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))
	err := os.Setenv("AL_SSM_ALL_MONGO_HOST", "")
	Expect(err).To(BeNil())
	v, err := servicessm.GetParameter(logger.NewEmptyLogger(), context.TODO(), "all", "/mongo/host", false)

	Expect(*v).To(Equal("asia-loop-admin.yncuk.mongodb.net"))
}
