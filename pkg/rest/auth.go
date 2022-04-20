package rest

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicessm"
)

const (
	HeaderAPIKey = "X-Al-Api-Key"
)

func AuthorizeAnalyticsRequest(log *logrus.Entry, ctx context.Context, request *events.APIGatewayProxyRequest) error {
	log.Infof("authorize analytics request [%s] [%s]", request.HTTPMethod, request.Path)

	keyInReq, hasKey := request.Headers[HeaderAPIKey]
	if !hasKey {
		return errors.New("missing API key")
	}

	key, err := servicessm.GetParameter(log, ctx, "all", "/analytics/key", true)
	if err != nil {
		return err
	}

	if keyInReq != *key {
		return errors.New("unauthorized")
	}

	return nil
}
