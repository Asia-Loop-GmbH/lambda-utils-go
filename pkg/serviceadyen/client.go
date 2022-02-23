package serviceadyen

import (
	"context"
	"fmt"

	"github.com/adyen/adyen-go-api-library/v5/src/adyen"
	"github.com/adyen/adyen-go-api-library/v5/src/common"
	"github.com/sirupsen/logrus"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicessm"
)

var (
	envMap = map[string]common.Environment{
		"dev":  common.TestEnv,
		"pre":  common.LiveEnv,
		"prod": common.LiveEnv,
	}
)

func newClient(log *logrus.Entry, ctx context.Context, stage string) (*adyen.APIClient, error) {
	log.Infof("new adyen client: %s", stage)
	apiKey, err := servicessm.GetParameter(log, ctx, stage, "/adyen/key", true)
	if err != nil {
		return nil, err
	}
	environment, ok := envMap[stage]
	if !ok {
		return nil, fmt.Errorf("no adyen environment config found for stage: %s", stage)
	}
	if stage == "dev" {
		return adyen.NewClient(&common.Config{
			ApiKey:      *apiKey,
			Environment: environment,
		}), nil
	}

	return adyen.NewClient(&common.Config{
		ApiKey:                *apiKey,
		Environment:           environment,
		LiveEndpointURLPrefix: livePrefix,
	}), nil
}
