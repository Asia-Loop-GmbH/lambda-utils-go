package dbadmin

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Increment struct {
	ID        primitive.ObjectID `bson:"_id"`
	Key       string             `bson:"key"`
	Value     int64              `bson:"value"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}
