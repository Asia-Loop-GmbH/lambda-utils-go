package address_test

import (
	. "github.com/onsi/gomega"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/address"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/logger"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/test"

	"testing"
)

func TestResolveAddress(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	Expect(address.ResolveAddress(logger.NewEmptyLogger(), "rudolf-schwarz platz 1 frankfurt")).To(Equal(&address.ResolveAddressResult{
		StreetNumber:     "1",
		Street:           "Rudolf-Schwarz-Platz",
		City:             "Frankfurt am Main",
		Postcode:         "60438",
		State:            "Hessen",
		FormattedAddress: "Rudolf-Schwarz-Platz 1, 60438 Frankfurt am Main, Deutschland",
	}))
}
