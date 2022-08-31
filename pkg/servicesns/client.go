package servicesns

import (
	"context"
	"sync"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

var (
	initClient sync.Once
	client     *sns.Client
)

func getClient(ctx context.Context) (*sns.Client, error) {
	var err error
	initClient.Do(func() {
		cfg, e := config.LoadDefaultConfig(ctx)
		if e != nil {
			err = e
			return
		}
		client = sns.NewFromConfig(cfg)
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}
