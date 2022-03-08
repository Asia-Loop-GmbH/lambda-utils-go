package servicedynamodb

import (
	"context"
	"sync"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/sirupsen/logrus"
)

var (
	initClient sync.Once
	client     *dynamodb.Client
)

func NewClient(log *logrus.Entry, ctx context.Context) (*dynamodb.Client, error) {
	var err error
	initClient.Do(func() {
		cfg, e := config.LoadDefaultConfig(ctx)
		if e != nil {
			err = e
			return
		}
		client = dynamodb.NewFromConfig(cfg)
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}
