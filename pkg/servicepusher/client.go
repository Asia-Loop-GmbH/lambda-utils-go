package servicepusher

import (
	"fmt"
	"sync"

	"github.com/pusher/pusher-http-go/v5"
	"github.com/sirupsen/logrus"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/myaws"
)

const (
	clusterEU = "eu"
)

var (
	initClient sync.Once
	client     *pusher.Client
)

func getClient(log *logrus.Entry, stage string) (*pusher.Client, error) {
	initClient.Do(func() {
		log.Infof("init pusher in [%s]", stage)
		app, err := myaws.GetSSMParameter(log, stage, "/pusher/app", false)
		if err != nil {
			return
		}
		key, err := myaws.GetSSMParameter(log, stage, "/pusher/key", false)
		if err != nil {
			return
		}
		secret, err := myaws.GetSSMParameter(log, stage, "/pusher/secret", true)
		if err != nil {
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
	if client == nil {
		log.Errorf("failed to initialize pusher in [%s]", stage)
		return nil, fmt.Errorf("failed to initialize pusher in [%s]", stage)
	}
	return client, nil
}
