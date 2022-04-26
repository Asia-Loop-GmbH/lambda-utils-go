package servicemailjet_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/internal/pkg/test"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/logger"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicemailjet"
)

func TestSendOrder(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	err := servicemailjet.SendOrder(logger.NewEmptyLogger(), context.TODO(), servicemailjet.SendInput{
		From: servicemailjet.Email{
			Address: "noreply@asialoop.de",
			Name:    "Asia Loop GmbH",
		},
		To: []servicemailjet.Email{{
			Address: "lenamtruong@gmail.com",
			Name:    "Nam Truong Le",
		}},
		Subject:    "Neue Bestellung #1234",
		TemplateID: servicemailjet.TemplateIDOrder,
	}, servicemailjet.SendOrderVariables{
		FirstName:     "Nam Truong",
		Title:         "vielen Dank für Deine Bestellung",
		Content:       "Wir haben Deine Bestellung empfangen. Bei Fragen oder Anmerkungen zu Deiner Bestellung möchten wir Dich bitten, das Restaurant telefonisch zu kontaktieren. Für weniger dringende Fragen kannst Du mit unserem Kundenservice Kontakt aufnehmen.",
		ActionText:    "Verfolge Deine Bestellung",
		ActionLink:    "https://google.de",
		ActionEnabled: "true",
	})

	Expect(err).To(BeNil())
}
