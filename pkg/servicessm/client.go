package servicessm

import (
	"context"
	"sync"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

var (
	initClient sync.Once
	client     *ssm.Client
)

func getClient(ctx context.Context) (*ssm.Client, error) {
	var err error
	initClient.Do(func() {
		cfg, e := config.LoadDefaultConfig(ctx)
		if e != nil {
			err = e
			return
		}
		client = ssm.NewFromConfig(cfg)
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}
