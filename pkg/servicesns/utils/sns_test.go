package utils_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/internal/pkg/test"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/logger"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicesns/utils"
)

func TestGetSNSStringAttribute_Success(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))
	v, err := utils.GetSNSStringAttribute(logger.NewEmptyLogger(), map[string]interface{}{"Value": "foo"})
	Expect(err).To(BeNil())
	Expect(v).To(Equal("foo"))
}

func TestGetSNSStringAttribute_NotString(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))
	_, err := utils.GetSNSStringAttribute(logger.NewEmptyLogger(), map[string]interface{}{"Value": 10})
	Expect(err).ToNot(BeNil())
}

func TestGetSNSStringAttribute_NoValue(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))
	_, err := utils.GetSNSStringAttribute(logger.NewEmptyLogger(), map[string]interface{}{})
	Expect(err).ToNot(BeNil())
}
