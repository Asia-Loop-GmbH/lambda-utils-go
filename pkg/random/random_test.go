package random_test

import (
	"log"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/internal/pkg/test"
	utils "github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/random"
)

func TestRandomString_Success(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	s := utils.String(10, true, true, true)
	log.Printf(s)
	Expect(len(s)).To(Equal(10))
}

func TestRandomString_Different(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	exists := map[string]bool{}

	for i := 0; i < 1000000; i++ {
		s := utils.String(10, true, true, true)
		_, ok := exists[s]
		Expect(ok).To(BeFalse())
		exists[s] = true
	}
}

func TestRandomString_Different_OrderID(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	exists := map[string]bool{}

	for i := 0; i < 1000000; i++ {
		s := utils.String(8, false, true, true)
		_, ok := exists[s]
		Expect(ok).To(BeFalse())
		exists[s] = true
	}
}
