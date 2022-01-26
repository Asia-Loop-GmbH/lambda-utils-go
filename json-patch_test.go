package lambda_utils_go_test

import (
	utils "github.com/asia-loop-gmbh/lambda-utils-go"
	"github.com/asia-loop-gmbh/lambda-utils-go/test"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

func TestMongoUpdateFromJSONPatch(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

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
	patches := []utils.JSONPatch{patch1, patch2, patch3}
	now := time.Now()
	update, err := utils.MongoUpdateFromJSONPatch(&patches, &now)

	Expect(err).To(BeNil())
	Expect(update).To(Equal(bson.A{
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
				"updatedAt": primitive.NewDateTimeFromTime(now),
			},
		},
	}))
}
