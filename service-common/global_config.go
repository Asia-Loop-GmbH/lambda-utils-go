package service_common

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/asia-loop-gmbh/lambda-types-go/admin"
	"github.com/asia-loop-gmbh/lambda-utils-go/mymongo"
)

func GetGlobalConfig(log *logrus.Entry, stage string) (*admin.GlobalConfig, error) {
	client, database, err := mymongo.NewMongoAdminClient(log, context.Background(), stage)
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
