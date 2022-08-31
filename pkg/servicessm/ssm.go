package servicessm

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	commoncontext "github.com/nam-truong-le/lambda-utils-go/pkg/context"
	"github.com/nam-truong-le/lambda-utils-go/pkg/logger"
	"github.com/pkg/errors"
)

// GetStageParameter returns ssm parameter. Stage is read from context.
func GetStageParameter(ctx context.Context, name string, decryption bool) (*string, error) {
	stage, ok := ctx.Value(commoncontext.FieldStage).(string)
	if !ok {
		return nil, fmt.Errorf("undefined stage in context")
	}

	return getParameterFromStage(ctx, stage, name, decryption)
}

// GetGlobalParameter returns ssm parameter for stage "all"
func GetGlobalParameter(ctx context.Context, name string, decryption bool) (*string, error) {
	return getParameterFromStage(ctx, "all", name, decryption)
}

func getParameterFromStage(ctx context.Context, stage, name string, decryption bool) (*string, error) {
	log := logger.FromContext(ctx)

	log.Infof("get [%s %s] variable", stage, name)
	if fromEnvVar := getFromEnvVar(stage, name); fromEnvVar != "" {
		log.Infof("found [%s %s] from env variable", stage, name)
		return &fromEnvVar, nil
	}
	c, err := getClient(ctx)
	if err != nil {
		return nil, err
	}
	parameterName := aws.String("/" + stage + name)
	getParameter, err := c.GetParameter(ctx, &ssm.GetParameterInput{
		Name:           parameterName,
		WithDecryption: decryption,
	})
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("could not find ssm parameter: %s", *parameterName))
	}
	log.Infof("found [%s %s] from aws ssm", stage, name)
	return getParameter.Parameter.Value, nil
}

func getFromEnvVar(stage, name string) string {
	name = strings.ReplaceAll(name, "/", "_")
	envKey := fmt.Sprintf("AL_SSM_%s%s", strings.ToUpper(stage), strings.ToUpper(name))
	return os.Getenv(envKey)
}
