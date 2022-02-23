package servicepusher

import (
	"context"
	"sync"

	"github.com/pkg/errors"
	"github.com/pusher/pusher-http-go/v5"
	"github.com/sirupsen/logrus"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicessm"
)

const (
	clusterEU = "eu"
)

var (
	initClient sync.Once
	client     *pusher.Client
)

func getClient(log *logrus.Entry, ctx context.Context, stage string) (*pusher.Client, error) {
	var err error
	initClient.Do(func() {
		log.Infof("init pusher in [%s]", stage)
		app, e := servicessm.GetParameter(log, ctx, stage, "/pusher/app", false)
		if e != nil {
			err = e
			return
		}
		key, e := servicessm.GetParameter(log, ctx, stage, "/pusher/key", false)
		if e != nil {
			err = e
			return
		}
		secret, e := servicessm.GetParameter(log, ctx, stage, "/pusher/secret", true)
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
		log.Infof("pusher initialized in [%s]", stage)
	})
	if err != nil {
		log.Errorf("failed to initialize pusher in [%s]: %s", stage, err)
		return nil, errors.Wrapf(err, "failed to initialize pusher in [%s]", stage)
	}
	return client, nil
}
