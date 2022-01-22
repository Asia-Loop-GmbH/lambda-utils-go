package lambda_utils_go

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"time"
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

func MongoUpdateFromJSONPatch(patches *[]JSONPatch, now *time.Time) (bson.A, error) {
	result := bson.A{}
	for _, patch := range *patches {
		singleUpdate, err := updateFromOnePatch(&patch)
		if err != nil {
			return nil, err
		}
		result = append(result, singleUpdate)
	}
	result = append(result, updateUpdatedAt(now))
	return result, nil
}

func updateFromOnePatch(patch *JSONPatch) (bson.M, error) {
	switch patch.OP {
	case jsonPatchOperatorReplace:
		return bson.M{
			mongoMethodSet: bson.M{
				jsonPathToMongoPath(&patch.Path): patch.Value,
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
