package service_common

import (
	"context"
	"github.com/asia-loop-gmbh/lambda-types-go/admin"
	"github.com/asia-loop-gmbh/lambda-utils-go/mongo"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func GetGlobalConfig(log *logrus.Entry, stage string) (*admin.GlobalConfig, error) {
	client, database, err := mongo.NewMongoAdminClient(log, context.Background(), stage)
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())

	log.Infof("database connected: %s", *database)
	collectionGlobalConfig := client.Database(*database).Collection(admin.CollectionGlobalConfig)
	find := collectionGlobalConfig.FindOne(context.Background(), bson.M{})
	globalConfig := new(admin.GlobalConfig)
	err = find.Decode(globalConfig)
	if err != nil {
		return nil, err
	}
	return globalConfig, nil
}
