package lambda_utils_go_test

import (
	utils "github.com/asia-loop-gmbh/lambda-utils-go"
	"github.com/asia-loop-gmbh/lambda-utils-go/test"
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
