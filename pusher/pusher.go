package pusher

import (
	"github.com/asia-loop-gmbh/lambda-utils-go/aws"
	"github.com/pusher/pusher-http-go/v5"
)

func NewPusherClient(stage *string) (*pusher.Client, error) {
	app, err := aws.GetSSMParameter(*stage, "/pusher/app", false)
	if err != nil {
		return nil, err
	}
	key, err := aws.GetSSMParameter(*stage, "/pusher/key", false)
	if err != nil {
		return nil, err
	}
	secret, err := aws.GetSSMParameter(*stage, "/pusher/secret", true)
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
