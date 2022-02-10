package pusher

import (
	"github.com/pusher/pusher-http-go/v5"
	"github.com/sirupsen/logrus"

	"github.com/asia-loop-gmbh/lambda-utils-go/myaws"
)

func NewPusherClient(log *logrus.Entry, stage *string) (*pusher.Client, error) {
	app, err := myaws.GetSSMParameter(log, *stage, "/pusher/app", false)
	if err != nil {
		return nil, err
	}
	key, err := myaws.GetSSMParameter(log, *stage, "/pusher/key", false)
	if err != nil {
		return nil, err
	}
	secret, err := myaws.GetSSMParameter(log, *stage, "/pusher/secret", true)
	if err != nil {
		return nil, err
	}

	client := pusher.Client{
		AppID:   *app,
		Key:     *key,
		Secret:  *secret,
		Cluster: "eu",
		Secure:  true,
	}

	return &client, nil
}
