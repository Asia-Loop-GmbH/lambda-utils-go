package servicemongo

import (
	"context"
	"time"

	"github.com/asia-loop-gmbh/lambda-types-go/v2/pkg/admin"
	"github.com/asia-loop-gmbh/lambda-utils-go/v4/pkg/dbadmin"
	"github.com/nam-truong-le/lambda-utils-go/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Next(ctx context.Context, key string) (int64, error) {
	log := logger.FromContext(ctx)
	log.Infof("next increment of [%s]", key)
	colIncrement, err := AdminCollection(ctx, dbadmin.CollectionIncrement)
	if err != nil {
		return 0, err
	}

	inc := new(admin.Increment)
	findExisting := colIncrement.FindOne(ctx, bson.M{"key": key})
	if err := findExisting.Decode(inc); err != nil {
		if err != mongo.ErrNoDocuments {
			log.Info("unwanted error from mongo")
			return 0, err
		}
		log.Infof("no existing increment found for [%s], create new increment", key)
		create, err := colIncrement.InsertOne(ctx, admin.Increment{
			ID:        primitive.NewObjectID(),
			Value:     0,
			Key:       key,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
		if err != nil {
			return 0, err
		}
		findNew := colIncrement.FindOne(ctx, bson.M{"_id": create.InsertedID})
		if err := findNew.Decode(inc); err != nil {
			return 0, err
		}
	}

	returnAfter := options.After
	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &returnAfter,
	}
	update := colIncrement.FindOneAndUpdate(ctx, bson.M{"key": key}, bson.M{
		"$set": bson.M{
			"updatedAt": time.Now(),
		},
		"$inc": bson.M{
			"value": 1,
		},
	}, &opts)
	updatedInc := new(admin.Increment)
	if err := update.Decode(updatedInc); err != nil {
		return 0, err
	}
	return updatedInc.Value, nil
}
