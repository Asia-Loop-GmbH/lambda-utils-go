package normalizer_test

import (
	"github.com/asia-loop-gmbh/lambda-utils-go/normalizer"
	"github.com/asia-loop-gmbh/lambda-utils-go/test"
	. "github.com/onsi/gomega"
	"testing"
)

func TestEmail(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	Expect(normalizer.Email("LeNamtrUong@gmail.com")).To(Equal("lenamtruong@gmail.com"))
}

func TestPhoneNumber(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	Expect(normalizer.PhoneNumber("1701234567")).To(Equal("+49 170 1234567"))
}

func TestName(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	Expect(normalizer.Name("  le     nam-truong     nhung  ")).To(Equal("Le Nam-Truong Nhung"))
}
