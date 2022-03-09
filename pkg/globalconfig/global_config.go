package globalconfig

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/dbadmin"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicemongo"
)

func GetGlobalConfig(log *logrus.Entry, ctx context.Context, stage string) (*dbadmin.GlobalConfig, error) {
	colGlbCfg, err := servicemongo.AdminCollection(log, ctx, stage, dbadmin.CollectionGlobalConfig)
	if err != nil {
		return nil, err
	}
	find := colGlbCfg.FindOne(context.Background(), bson.M{})
	globalConfig := new(dbadmin.GlobalConfig)
	err = find.Decode(globalConfig)
	if err != nil {
		return nil, err
	}
	return globalConfig, nil
}
