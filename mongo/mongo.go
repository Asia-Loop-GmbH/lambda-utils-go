package mongo

import (
	"context"
	"fmt"
	"github.com/asia-loop-gmbh/lambda-utils-go/aws"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoCorpClient(logger logrus.Entry, ctx context.Context, env string) (*mongo.Client, *string, error) {
	return newMongoClient(logger, ctx, env, "corp")
}

func NewMongoAdminClient(logger logrus.Entry, ctx context.Context, env string) (*mongo.Client, *string, error) {
	return newMongoClient(logger, ctx, env, "admin")
}

func newMongoClient(logger logrus.Entry, ctx context.Context, env string, app string) (*mongo.Client, *string, error) {

	mongoHost, err := aws.GetSSMParameter("all", "/mongo/host", false)
	if err != nil {
		return nil, nil, err
	}
	mongoUsername, err := aws.GetSSMParameter(env, fmt.Sprintf("/%s/mongo/username", app), false)
	if err != nil {
		return nil, nil, err
	}
	mongoPassword, err := aws.GetSSMParameter(env, fmt.Sprintf("/%s/mongo/password", app), true)
	if err != nil {
		return nil, nil, err
	}
	mongoDatabase, err := aws.GetSSMParameter(env, fmt.Sprintf("/%s/mongo/database", app), false)
	if err != nil {
		return nil, nil, err
	}

	logger.Info("mongo host = %s, mongo db = %s", *mongoHost, *mongoDatabase)

	mongoFullUrl := fmt.Sprintf("mongodb+srv://%s", *mongoHost)

	client, err := mongo.NewClient(
		options.Client().ApplyURI(mongoFullUrl).SetAuth(options.Credential{
			Username: *mongoUsername,
			Password: *mongoPassword,
		}),
	)
	if err != nil {
		return nil, nil, err
	}
	err = client.Connect(ctx)
	if err != nil {
		return nil, nil, err
	}

	return client, mongoDatabase, nil
}
