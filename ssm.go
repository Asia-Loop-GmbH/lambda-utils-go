package lambda_utils_go

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func GetSSMParameter(env, name string, decryption bool) (*string, error) {
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
	return getParameter.Parameter.Value, nil
}
