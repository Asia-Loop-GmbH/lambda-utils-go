package myaws

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func GetSSMParameter(log *logrus.Entry, stage, name string, decryption bool) (*string, error) {
	log.Infof("get [%s %s] variable", stage, name)
	if fromEnvVar, ok := getFromEnvVar(stage, name); ok {
		log.Infof("found [%s %s] from env variable", stage, name)
		return &fromEnvVar, nil
	}
	awsSession, err := session.NewSession(&aws.Config{Region: aws.String(endpoints.EuCentral1RegionID)})
	if err != nil {
		return nil, err
	}
	ssmService := ssm.New(awsSession)
	parameterName := aws.String("/" + stage + name)
	getParameter, err := ssmService.GetParameter(&ssm.GetParameterInput{
		Name:           parameterName,
		WithDecryption: aws.Bool(decryption),
	})
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("could not find ssm parameter: %s", *parameterName))
	}
	log.Infof("found [%s %s] from aws ssm", stage, name)
	return getParameter.Parameter.Value, nil
}

func getFromEnvVar(stage, name string) (string, bool) {
	name = strings.ReplaceAll(name, "/", "_")
	envKey := fmt.Sprintf("AL_SSM_%s%s", strings.ToUpper(stage), strings.ToUpper(name))
	return os.LookupEnv(envKey)
}
