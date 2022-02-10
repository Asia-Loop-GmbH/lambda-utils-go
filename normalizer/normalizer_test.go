package normalizer_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/asia-loop-gmbh/lambda-utils-go/logger"
	"github.com/asia-loop-gmbh/lambda-utils-go/normalizer"
	"github.com/asia-loop-gmbh/lambda-utils-go/test"
)

func TestEmail(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	Expect(normalizer.Email(logger.NewEmptyLogger(), "LeNamtrUong@gmail.com")).To(Equal("lenamtruong@gmail.com"))
}

func TestPhoneNumber(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	Expect(normalizer.PhoneNumber(logger.NewEmptyLogger(), "1701234567")).To(Equal("+49 170 1234567"))
}

func TestName(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	Expect(normalizer.Name(logger.NewEmptyLogger(), "  le     nam-truong     nhung  ")).To(Equal("Le Nam-Truong Nhung"))
}
