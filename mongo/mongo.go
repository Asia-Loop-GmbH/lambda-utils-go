package mongo

import (
	"context"
	"fmt"
	"github.com/asia-loop-gmbh/lambda-utils-go/aws"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func NewMongoAdminClient(ctx context.Context, env string) (*mongo.Client, *string, error) {

	mongoHost, err := aws.GetSSMParameter("all", "/mongo/host", false)
	if err != nil {
		return nil, nil, err
	}
	mongoUsername, err := aws.GetSSMParameter(env, "/admin/mongo/username", false)
	if err != nil {
		return nil, nil, err
	}
	mongoPassword, err := aws.GetSSMParameter(env, "/admin/mongo/password", true)
	if err != nil {
		return nil, nil, err
	}
	mongoDatabase, err := aws.GetSSMParameter(env, "/admin/mongo/database", false)
	if err != nil {
		return nil, nil, err
	}

	log.Printf("mongo host = %s, mongo db = %s", *mongoHost, *mongoDatabase)

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
