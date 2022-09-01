package servicemailjet_test

import (
	"context"
	"testing"

	commoncontext "github.com/nam-truong-le/lambda-utils-go/pkg/context"
	"github.com/stretchr/testify/assert"

	"github.com/asia-loop-gmbh/lambda-utils-go/v4/pkg/servicemailjet"
)

func TestSend(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	ctx := context.WithValue(context.TODO(), commoncontext.FieldStage, "dev")
	err := servicemailjet.Send(ctx, servicemailjet.SendInput{
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
	}, map[string]interface{}{
		"firstName":     "Nam Truong",
		"title":         "vielen Dank für Deine Bestellung",
		"content":       "Wir haben Deine Bestellung empfangen. Bei Fragen oder Anmerkungen zu Deiner Bestellung möchten wir Dich bitten, das Restaurant telefonisch zu kontaktieren. Für weniger dringende Fragen kannst Du mit unserem Kundenservice Kontakt aufnehmen.",
		"actionText":    "Verfolge Deine Bestellung",
		"actionLink":    "https://google.de",
		"actionEnabled": "true",
	})

	assert.NoError(t, err)
}
