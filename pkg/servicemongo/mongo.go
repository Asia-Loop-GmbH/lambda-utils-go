package servicemongo

import (
	"context"
	"fmt"
	"sync"

	"github.com/asia-loop-gmbh/lambda-utils-go/v4/pkg/servicessm"
	commoncontext "github.com/nam-truong-le/lambda-utils-go/pkg/context"
	"github.com/nam-truong-le/lambda-utils-go/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func Disconnect(ctx context.Context) {
	log := logger.FromContext(ctx)
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

func CorpCollection(ctx context.Context, collection string) (*mongo.Collection, error) {
	return getCollection(ctx, "corp", collection)
}

func AdminCollection(ctx context.Context, collection string) (*mongo.Collection, error) {
	return getCollection(ctx, "admin", collection)
}

func getCollection(ctx context.Context, app, collection string) (*mongo.Collection, error) {
	log := logger.FromContext(ctx)

	stage, ok := ctx.Value(commoncontext.FieldStage).(string)
	if !ok {
		return nil, fmt.Errorf("undefined stage in context")
	}

	client, ok := clients[app]
	if !ok {
		return nil, fmt.Errorf("no client registered for [%s]", app)
	}

	client.Init.Do(func() {
		log.Infof("first time init")
		mongoHost, err := servicessm.GetGlobalParameter(ctx, "/mongo/host", false)
		if err != nil {
			log.Errorf("failed: %s", err)
			return
		}
		mongoUsername, err := servicessm.GetStageParameter(ctx, fmt.Sprintf("/%s/mongo/username", app), false)
		if err != nil {
			log.Errorf("failed: %s", err)
			return
		}
		mongoPassword, err := servicessm.GetStageParameter(ctx, fmt.Sprintf("/%s/mongo/password", app), true)
		if err != nil {
			log.Errorf("failed: %s", err)
			return
		}
		mongoDatabase, err := servicessm.GetStageParameter(ctx, fmt.Sprintf("/%s/mongo/database", app), false)
		if err != nil {
			log.Errorf("failed: %s", err)
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
			log.Errorf("failed: %s", err)
			return
		}
		err = c.Connect(ctx)
		if err != nil {
			log.Errorf("failed: %s", err)
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
