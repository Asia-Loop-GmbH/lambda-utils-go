package servicemongo

import (
	"context"
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicessm"
)

type dbClient struct {
	Init     sync.Once
	Client   *mongo.Client
	Database *string
}

var clients = map[string]*dbClient{
	"admin": {
		Init:     sync.Once{},
		Client:   nil,
		Database: nil,
	},
	"corp": {
		Init:     sync.Once{},
		Client:   nil,
		Database: nil,
	},
}

func Disconnect(log *logrus.Entry, ctx context.Context) {
	for k, c := range clients {
		if c.Client != nil {
			log.Infof("disconnect client [%s]", k)
			err := c.Client.Disconnect(ctx)
			if err != nil {
				log.Errorf("failed to disconnect client [%s]: %s", k, err)
			} else {
				log.Infof("client [%s] disconnected", k)
			}
			c.Init = sync.Once{}
			c.Client = nil
			c.Database = nil
			log.Infof("client [%s] reset", k)
		}
	}
}

func CorpCollection(log *logrus.Entry, ctx context.Context, stage, collection string) (*mongo.Collection, error) {
	return getCollection(log, ctx, stage, "corp", collection)
}

func AdminCollection(log *logrus.Entry, ctx context.Context, stage, collection string) (*mongo.Collection, error) {
	return getCollection(log, ctx, stage, "admin", collection)
}

func getCollection(log *logrus.Entry, ctx context.Context, stage, app, collection string) (*mongo.Collection, error) {

	client, ok := clients[app]
	if !ok {
		return nil, fmt.Errorf("no client registered for [%s]", app)
	}

	client.Init.Do(func() {
		log.Infof("first time init")
		mongoHost, err := servicessm.GetParameter(log, ctx, "all", "/mongo/host", false)
		if err != nil {
			return
		}
		mongoUsername, err := servicessm.GetParameter(log, ctx, stage, fmt.Sprintf("/%s/mongo/username", app), false)
		if err != nil {
			return
		}
		mongoPassword, err := servicessm.GetParameter(log, ctx, stage, fmt.Sprintf("/%s/mongo/password", app), true)
		if err != nil {
			return
		}
		mongoDatabase, err := servicessm.GetParameter(log, ctx, stage, fmt.Sprintf("/%s/mongo/database", app), false)
		if err != nil {
			return
		}
		log.Infof("host = %s, user = %s, db = %s", *mongoHost, *mongoUsername, *mongoDatabase)
		mongoFullUrl := fmt.Sprintf("mongodb+srv://%s", *mongoHost)

		c, err := mongo.NewClient(
			options.Client().ApplyURI(mongoFullUrl).SetAuth(options.Credential{
				Username: *mongoUsername,
				Password: *mongoPassword,
			}),
		)
		if err != nil {
			return
		}
		err = c.Connect(ctx)
		if err != nil {
			return
		}

		client.Client = c
		client.Database = mongoDatabase
	})

	if client.Client == nil {
		return nil, fmt.Errorf("could not init db: %s, %s", stage, app)
	}
	return client.Client.Database(*client.Database).Collection(collection), nil
}
