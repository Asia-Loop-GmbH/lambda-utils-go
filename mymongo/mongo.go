package mymongo

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/asia-loop-gmbh/lambda-utils-go/myaws"
)

func NewMongoCorpClient(log *logrus.Entry, ctx context.Context, stage string) (*mongo.Client, *string, error) {
	return newMongoClient(log, ctx, stage, "corp")
}

func NewMongoAdminClient(log *logrus.Entry, ctx context.Context, stage string) (*mongo.Client, *string, error) {
	return newMongoClient(log, ctx, stage, "admin")
}

func newMongoClient(log *logrus.Entry, ctx context.Context, stage string, app string) (*mongo.Client, *string, error) {

	mongoHost, err := myaws.GetSSMParameter(log, "all", "/mongo/host", false)
	if err != nil {
		return nil, nil, err
	}
	mongoUsername, err := myaws.GetSSMParameter(log, stage, fmt.Sprintf("/%s/mongo/username", app), false)
	if err != nil {
		return nil, nil, err
	}
	mongoPassword, err := myaws.GetSSMParameter(log, stage, fmt.Sprintf("/%s/mongo/password", app), true)
	if err != nil {
		return nil, nil, err
	}
	mongoDatabase, err := myaws.GetSSMParameter(log, stage, fmt.Sprintf("/%s/mongo/database", app), false)
	if err != nil {
		return nil, nil, err
	}

	log.Infof("mongo host = %s, mongo db = %s", *mongoHost, *mongoDatabase)

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
