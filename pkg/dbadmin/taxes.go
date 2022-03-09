package dbadmin

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tax struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	WPID      int                `bson:"id" json:"wpId"`
	Rate      string             `bson:"rate" json:"rate"`
	Name      string             `bson:"name" json:"name"`
	TaxClass  string             `bson:"class" json:"taxClass"` // it's ok to have different names here because we don't provide PATCH request for this entity.
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}
