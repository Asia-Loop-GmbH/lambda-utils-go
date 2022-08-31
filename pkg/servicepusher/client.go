package servicepusher

import (
	"context"
	"sync"

	"github.com/asia-loop-gmbh/lambda-utils-go/v4/pkg/servicessm"
	"github.com/nam-truong-le/lambda-utils-go/pkg/logger"
	"github.com/pkg/errors"
	"github.com/pusher/pusher-http-go/v5"
)

const (
	clusterEU = "eu"
)

var (
	initClient sync.Once
	client     *pusher.Client
)

func getClient(ctx context.Context) (*pusher.Client, error) {
	log := logger.FromContext(ctx)
	var err error
	initClient.Do(func() {
		log.Infof("init pusher")
		app, e := servicessm.GetStageParameter(ctx, "/pusher/app", false)
		if e != nil {
			err = e
			return
		}
		key, e := servicessm.GetStageParameter(ctx, "/pusher/key", false)
		if e != nil {
			err = e
			return
		}
		secret, e := servicessm.GetStageParameter(ctx, "/pusher/secret", true)
		if e != nil {
			err = e
			return
		}
		client = &pusher.Client{
			AppID:   *app,
			Key:     *key,
			Secret:  *secret,
			Cluster: clusterEU,
			Secure:  true,
		}
		log.Infof("pusher initialized")
	})
	if err != nil {
		log.Errorf("failed to initialize pusher: %s", err)
		return nil, errors.Wrapf(err, "failed to initialize pusher")
	}
	return client, nil
}
