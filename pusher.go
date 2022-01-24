package lambda_utils_go

import "github.com/pusher/pusher-http-go/v5"

func NewPusherClient(stage *string) (*pusher.Client, error) {
	app, err := GetSSMParameter(*stage, "/pusher/app", false)
	if err != nil {
		return nil, err
	}
	key, err := GetSSMParameter(*stage, "/pusher/key", false)
	if err != nil {
		return nil, err
	}
	secret, err := GetSSMParameter(*stage, "/pusher/secret", true)
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
