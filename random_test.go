package lambda_utils_go_test

import (
	utils "github.com/asia-loop-gmbh/lambda-utils-go"
	. "github.com/onsi/gomega"
	"log"
	"testing"
)

func TestRandomStringSuccess(t *testing.T) {
	RegisterFailHandler(failedHandler(t))

	s := utils.RandomString(10, true, true, true)
	log.Printf(s)
	Expect(len(s)).To(Equal(10))
}
