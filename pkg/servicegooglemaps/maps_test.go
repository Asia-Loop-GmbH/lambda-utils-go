package servicegooglemaps_test

import (
	"context"

	. "github.com/onsi/gomega"

	"testing"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/internal/pkg/test"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/logger"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicegooglemaps"
)

func TestResolveAddress(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	Expect(servicegooglemaps.ResolveAddress(logger.NewEmptyLogger(), context.TODO(), "rudolf-schwarz platz 1 frankfurt")).To(Equal(&servicegooglemaps.ResolveAddressResult{
		StreetNumber:     "1",
		Street:           "Rudolf-Schwarz-Platz",
		City:             "Frankfurt am Main",
		Postcode:         "60438",
		State:            "Hessen",
		FormattedAddress: "Rudolf-Schwarz-Platz 1, 60438 Frankfurt am Main, Deutschland",
	}))
}
