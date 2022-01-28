package address_test

import (
	"github.com/asia-loop-gmbh/lambda-utils-go/address"
	"github.com/asia-loop-gmbh/lambda-utils-go/test"
	. "github.com/onsi/gomega"
	"testing"
)

func TestResolveAddress(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	Expect(address.ResolveAddress("rudolf-schwarz platz 1 frankfurt")).To(Equal(&address.ResolveAddressResult{
		StreetNumber:     "1",
		Street:           "Rudolf-Schwarz-Platz",
		City:             "Frankfurt am Main",
		Postcode:         "60438",
		State:            "Hessen",
		FormattedAddress: "Rudolf-Schwarz-Platz 1, 60438 Frankfurt am Main, Deutschland",
	}))
}
