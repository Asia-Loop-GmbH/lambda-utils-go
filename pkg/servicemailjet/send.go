package servicemailjet

import (
	"context"

	"github.com/mailjet/mailjet-apiv3-go/v3"
	"github.com/sirupsen/logrus"
)

type TemplateID int

const (
	TemplateIDOrder TemplateID = 3889572
	ccEmail                    = "order@asialoop.de"
)

type Email struct {
	Address string
	Name    string
}

type SendInput struct {
	From       Email
	To         []Email
	Subject    string
	TemplateID TemplateID
}

func Send(log *logrus.Entry, ctx context.Context, input SendInput, variables map[string]interface{}) error {
	c, err := newClient(log, ctx)
	if err != nil {
		return err
	}

	receivers := make(mailjet.RecipientsV31, 0)
	for _, to := range input.To {
		receivers = append(receivers, mailjet.RecipientV31{
			Email: to.Address,
			Name:  to.Name,
		})
	}

	cc := make(mailjet.RecipientsV31, 0)
	cc = append(cc, mailjet.RecipientV31{Email: ccEmail})

	info := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: input.From.Address,
				Name:  input.From.Name,
			},
			To:               &receivers,
			Cc:               &cc,
			Subject:          input.Subject,
			TemplateID:       int(input.TemplateID),
			TemplateLanguage: true,
			Variables:        variables,
		},
	}

	_, err = c.SendMailV31(&mailjet.MessagesV31{Info: info})
	if err != nil {
		return err
	}
	return nil
}
