package rest

import (
	"context"

	"github.com/asia-loop-gmbh/lambda-utils-go/v4/pkg/servicessm"
	"github.com/aws/aws-lambda-go/events"
	"github.com/nam-truong-le/lambda-utils-go/pkg/logger"
	"github.com/pkg/errors"
)

const (
	HeaderAPIKey = "x-al-api-key"
)

func AuthorizeAnalyticsRequest(ctx context.Context, request *events.APIGatewayProxyRequest) error {
	log := logger.FromContext(ctx)

	log.Infof("authorize analytics request [%s] [%s]", request.HTTPMethod, request.Path)

	keyInReq, hasKey := request.Headers[HeaderAPIKey]
	if !hasKey {
		return errors.New("missing API key")
	}

	key, err := servicessm.GetGlobalParameter(ctx, "/analytics/key", true)
	if err != nil {
		return err
	}

	if keyInReq != *key {
		return errors.New("unauthorized")
	}

	return nil
}
