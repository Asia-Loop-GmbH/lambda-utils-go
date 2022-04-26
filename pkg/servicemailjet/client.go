package servicemailjet

import (
	"context"
	"sync"

	"github.com/mailjet/mailjet-apiv3-go/v3"
	"github.com/sirupsen/logrus"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicessm"
)

var (
	initClient sync.Once
	client     *mailjet.Client
)

func newClient(log *logrus.Entry, ctx context.Context) (*mailjet.Client, error) {
	var err error
	initClient.Do(func() {
		key, e := servicessm.GetParameter(log, ctx, "all", "/mailjet/key", false)
		if e != nil {
			err = e
			return
		}
		secret, e := servicessm.GetParameter(log, ctx, "all", "/mailjet/secret", true)
		if e != nil {
			err = e
			return
		}
		client = mailjet.NewMailjetClient(*key, *secret)
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}
