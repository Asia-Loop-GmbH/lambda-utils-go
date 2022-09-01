package servicemailjet

import (
	"context"
	"sync"

	"github.com/asia-loop-gmbh/lambda-utils-go/v4/pkg/servicessm"
	"github.com/mailjet/mailjet-apiv3-go/v3"
)

var (
	initClient sync.Once
	client     *mailjet.Client
)

func newClient(ctx context.Context) (*mailjet.Client, error) {
	var err error
	initClient.Do(func() {
		key, e := servicessm.GetGlobalParameter(ctx, "/mailjet/key", false)
		if e != nil {
			err = e
			return
		}
		secret, e := servicessm.GetGlobalParameter(ctx, "/mailjet/secret", true)
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
