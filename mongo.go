package lambda_utils_go

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func NewMongoAdminClient(ctx context.Context, env string) (*mongo.Client, error) {

	mongoHost, err := GetSSMParameter("all", "/mongo/host", false)
	if err != nil {
		return nil, err
	}
	mongoUsername, err := GetSSMParameter(env, "/admin/mongo/username", false)
	if err != nil {
		return nil, err
	}
	mongoPassword, err := GetSSMParameter(env, "/admin/mongo/password", true)
	if err != nil {
		return nil, err
	}
	mongoDatabase, err := GetSSMParameter(env, "/admin/mongo/database", false)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}
