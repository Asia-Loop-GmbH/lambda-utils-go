package dbadmin

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GlobalConfig struct {
	ID                           primitive.ObjectID `bson:"_id" json:"id"`
	ProductAttributeOutOfStockIn int                `bson:"productAttributeOutOfStockInId" json:"productAttributeOutOfStockInId"`
	ProductAttributePfandId      int                `bson:"productAttributePfandId" json:"productAttributePfandId"`
	PusherAPIKey                 string             `json:"pusherApiKey"` // TODO: before provide PATCH method, add this to DB!!!
	CreatedAt                    time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt                    time.Time          `bson:"updatedAt" json:"updatedAt"`
}
