package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"log"
	"os"
	"strings"
)

func GetSSMParameter(env, name string, decryption bool) (*string, error) {
	if fromEnvVar, ok := getFromEnVar(env, name); ok {
		log.Printf("get [%s %s] from environment variable", env, name)
		return &fromEnvVar, nil
	}
	awsSession, err := session.NewSession(&aws.Config{Region: aws.String("eu-central-1")})
	if err != nil {
		return nil, err
	}
	ssmService := ssm.New(awsSession)
	getParameter, err := ssmService.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String("/" + env + name),
		WithDecryption: aws.Bool(decryption),
	})
	if err != nil {
		return nil, err
	}
	log.Printf("get [%s %s] from aws ssm", env, name)
	return getParameter.Parameter.Value, nil
}

func getFromEnVar(env, name string) (string, bool) {
	name = strings.ReplaceAll(strings.ToUpper(name), "/", "_")
	envKey := fmt.Sprintf("AL_SSM_%s_%s", env, name)
	return os.LookupEnv(envKey)
}
