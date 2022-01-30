package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/pkg/errors"
	"log"
	"os"
	"strings"
)

func GetSSMParameter(env, name string, decryption bool) (*string, error) {
	if fromEnvVar, ok := getFromEnvVar(env, name); ok {
		log.Printf("get [%s %s] from environment variable", env, name)
		return &fromEnvVar, nil
	}
	awsSession, err := session.NewSession(&aws.Config{Region: aws.String("eu-central-1")})
	if err != nil {
		return nil, err
	}
	ssmService := ssm.New(awsSession)
	parameterName := aws.String("/" + env + name)
	getParameter, err := ssmService.GetParameter(&ssm.GetParameterInput{
		Name:           parameterName,
		WithDecryption: aws.Bool(decryption),
	})
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("could not find ssm parameter: %s", *parameterName))
	}
	log.Printf("get [%s %s] from aws ssm", env, name)
	return getParameter.Parameter.Value, nil
}

func getFromEnvVar(stage, name string) (string, bool) {
	name = strings.ReplaceAll(name, "/", "_")
	envKey := fmt.Sprintf("AL_SSM_%s%s", strings.ToUpper(stage), strings.ToUpper(name))
	return os.LookupEnv(envKey)
}
