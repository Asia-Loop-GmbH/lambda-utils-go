package servicecognito

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/sirupsen/logrus"
)

type DeleteUserData struct {
	Username string `json:"username"`
}

func DeleteUser(log *logrus.Entry, ctx context.Context, data *DeleteUserData) error {
	client, err := getClient(log, ctx)
	if err != nil {
		return err
	}
	_, err = client.AdminDeleteUser(ctx, &cognitoidentityprovider.AdminDeleteUserInput{
		UserPoolId: aws.String(cognitoPool),
		Username:   aws.String(data.Username),
	})
	if err != nil {
		return err
	}
	return nil
}
