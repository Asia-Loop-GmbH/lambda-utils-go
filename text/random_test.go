package text_test

import (
	"log"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/asia-loop-gmbh/lambda-utils-go/test"
	utils "github.com/asia-loop-gmbh/lambda-utils-go/text"
)

func TestRandomString_Success(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	s := utils.RandomString(10, true, true, true)
	log.Printf(s)
	Expect(len(s)).To(Equal(10))
}

func TestRandomString_Different(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	exists := map[string]bool{}

	for i := 0; i < 1000000; i++ {
		s := utils.RandomString(10, true, true, true)
		_, ok := exists[s]
		Expect(ok).To(BeFalse())
		exists[s] = true
	}
}

func TestRandomString_Different_OrderID(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	exists := map[string]bool{}

	for i := 0; i < 1000000; i++ {
		s := utils.RandomString(8, false, true, true)
		_, ok := exists[s]
		Expect(ok).To(BeFalse())
		exists[s] = true
	}
}
