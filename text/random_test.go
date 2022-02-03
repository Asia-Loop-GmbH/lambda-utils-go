package text_test

import (
	"github.com/asia-loop-gmbh/lambda-utils-go/test"
	utils "github.com/asia-loop-gmbh/lambda-utils-go/text"
	. "github.com/onsi/gomega"
	"log"
	"testing"
)

func TestRandomStringSuccess(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	s := utils.RandomString(10, true, true, true)
	log.Printf(s)
	Expect(len(s)).To(Equal(10))
}

func TestRandomStringDifferent(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	exists := map[string]bool{}

	for i := 0; i < 1000000; i++ {
		s := utils.RandomString(10, true, true, true)
		_, ok := exists[s]
		Expect(ok).To(BeFalse())
		exists[s] = true
	}
}
