package globalconfig

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/asia-loop-gmbh/lambda-utils-go/v4/pkg/dbadmin"
	"github.com/asia-loop-gmbh/lambda-utils-go/v4/pkg/servicemongo"
)

func GetGlobalConfig(ctx context.Context) (*dbadmin.GlobalConfig, error) {
	colGlbCfg, err := servicemongo.AdminCollection(ctx, dbadmin.CollectionGlobalConfig)
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
