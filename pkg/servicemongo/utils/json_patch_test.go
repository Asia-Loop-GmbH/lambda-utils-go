package utils_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/asia-loop-gmbh/lambda-utils-go/v4/pkg/servicemongo/utils"
)

func TestMongoUpdateFromJSONPatch(t *testing.T) {

	patch1 := utils.JSONPatch{
		OP:    "replace",
		Path:  "/a/b/c",
		Value: "string",
	}
	patch2 := utils.JSONPatch{
		OP:    "replace",
		Path:  "/a/b/d",
		Value: 1234,
	}
	patch3 := utils.JSONPatch{
		OP:    "replace",
		Path:  "/a/c/e",
		Value: true,
	}
	objectID1 := primitive.NewObjectID()
	objectID1String := objectID1.Hex()
	patchObjectID := utils.JSONPatch{
		OP:    "replace",
		Path:  "/objectId1/a",
		Value: objectID1String,
	}
	objectID2 := primitive.NewObjectID()
	objectID2String := objectID2.Hex()
	objectID3 := primitive.NewObjectID()
	objectID3String := objectID3.Hex()
	patchObjectIDArray := utils.JSONPatch{
		OP:    "replace",
		Path:  "/objectId1/array",
		Value: []interface{}{objectID2String, objectID3String},
	}
	patches := []utils.JSONPatch{patch1, patch2, patch3, patchObjectID, patchObjectIDArray}
	now := time.Now()
	objectIDPaths := []string{"objectId1.a", "objectId1.array"}
	update, err := utils.MongoUpdateFromJSONPatch(context.Background(), &patches, &now, objectIDPaths)

	assert.NoError(t, err)
	expected := bson.A{
		bson.M{
			"$set": bson.M{
				"a.b.c": "string",
			},
		},
		bson.M{
			"$set": bson.M{
				"a.b.d": 1234,
			},
		},
		bson.M{
			"$set": bson.M{
				"a.c.e": true,
			},
		},
		bson.M{
			"$set": bson.M{
				"objectId1.a": objectID1,
			},
		},
		bson.M{
			"$set": bson.M{
				"objectId1.array": []primitive.ObjectID{objectID2, objectID3},
			},
		},
		bson.M{
			"$set": bson.M{
				"updatedAt": primitive.NewDateTimeFromTime(now),
			},
		},
	}
	assert.Equal(t, expected, update)
}
