package utils

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/nam-truong-le/lambda-utils-go/pkg/logger"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JSONPatch struct {
	OP    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value"`
}

const (
	jsonPathSeparator        = "/"
	mongoPathSeparator       = "."
	jsonPatchOperatorReplace = "replace"
	mongoMethodSet           = "$set"
	mongoAttributeUpdatedAt  = "updatedAt"
)

func MongoUpdateFromJSONPatch(ctx context.Context, patches *[]JSONPatch, now *time.Time, objectIDPaths []string) (bson.A, error) {
	log := logger.FromContext(ctx)
	log.Infof("create update bson from json patch: %v", patches)
	result := bson.A{}
	for _, patch := range *patches {
		singleUpdate, err := updateFromOnePatch(&patch, objectIDPaths)
		if err != nil {
			return nil, err
		}
		result = append(result, singleUpdate)
	}
	result = append(result, updateUpdatedAt(now))
	log.Infof("JSON patch %v converted to mongo update: %v", patches, result)
	return result, nil
}

func updateFromOnePatch(patch *JSONPatch, objectIDPaths []string) (bson.M, error) {
	var err error
	mongoPath := jsonPathToMongoPath(&patch.Path)
	isObjectID := false
	for _, p := range objectIDPaths {
		if p == mongoPath {
			isObjectID = true
			break
		}
	}
	value := patch.Value
	if isObjectID {
		switch v := patch.Value.(type) {
		case string:
			value, err = primitive.ObjectIDFromHex(v)
			if err != nil {
				return nil, errors.Wrap(err, fmt.Sprintf("%s is not an uuid", v))
			}
		case []interface{}:
			uuids := make([]primitive.ObjectID, 0)
			for _, id := range v {
				switch vid := id.(type) {
				case string:
					uuid, err := primitive.ObjectIDFromHex(vid)
					if err != nil {
						return nil, errors.Wrap(err, fmt.Sprintf("%s is not an uuid", id))
					}
					uuids = append(uuids, uuid)
				default:
					return nil, fmt.Errorf("path %s is marked as uuid but got array element %v", mongoPath, vid)
				}
			}
			value = uuids
		default:
			return nil, fmt.Errorf("path %s is marked as uuid but got %v", mongoPath, v)
		}
	}
	switch patch.OP {
	case jsonPatchOperatorReplace:
		return bson.M{
			mongoMethodSet: bson.M{
				mongoPath: value,
			},
		}, nil
	default:
		return nil, fmt.Errorf("json patch operator '%s' not supported", patch.OP)
	}
}

func jsonPathToMongoPath(path *string) string {
	parts := strings.Split(*path, jsonPathSeparator)
	return strings.Join(parts[1:], mongoPathSeparator)
}

func updateUpdatedAt(now *time.Time) bson.M {
	return bson.M{
		mongoMethodSet: bson.M{
			mongoAttributeUpdatedAt: primitive.NewDateTimeFromTime(*now),
		},
	}
}
