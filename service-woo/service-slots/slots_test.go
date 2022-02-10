package service_slots_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/asia-loop-gmbh/lambda-utils-go/v2/logger"
	serviceslots "github.com/asia-loop-gmbh/lambda-utils-go/v2/service-woo/service-slots"
	"github.com/asia-loop-gmbh/lambda-utils-go/v2/test"
)

func TestGetSlots_Success(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	slots, err := serviceslots.GetSlots(logger.NewEmptyLogger(), "dev")

	Expect(err).To(BeNil())
	Expect(slots).NotTo(BeNil())
}
