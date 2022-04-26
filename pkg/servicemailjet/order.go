package servicemailjet

import (
	"context"

	"github.com/sirupsen/logrus"
)

type SendOrderVariables struct {
	FirstName     string
	Title         string
	Content       string
	ActionText    string
	ActionLink    string
	ActionEnabled string
}

func SendOrder(log *logrus.Entry, ctx context.Context, input SendInput, variables SendOrderVariables) error {
	return Send(log, ctx, input, map[string]interface{}{
		"firstName":     variables.FirstName,
		"title":         variables.Title,
		"content":       variables.Content,
		"actionText":    variables.ActionText,
		"actionLink":    variables.ActionLink,
		"actionEnabled": variables.ActionEnabled,
	})
}
