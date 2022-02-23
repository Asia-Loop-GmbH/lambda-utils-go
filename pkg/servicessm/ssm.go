package servicessm

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func GetParameter(log *logrus.Entry, ctx context.Context, stage, name string, decryption bool) (*string, error) {
	log.Infof("get [%s %s] variable", stage, name)
	if fromEnvVar := getFromEnvVar(stage, name); fromEnvVar != "" {
		log.Infof("found [%s %s] from env variable", stage, name)
		return &fromEnvVar, nil
	}
	client, err := getClient(log, ctx)
	if err != nil {
		return nil, err
	}
	parameterName := aws.String("/" + stage + name)
	getParameter, err := client.GetParameter(ctx, &ssm.GetParameterInput{
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
