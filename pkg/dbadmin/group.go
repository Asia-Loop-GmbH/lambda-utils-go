package dbadmin

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	GroupObjectIDs = []string{
		"orders", "driver", "store",
	}
)

type Group struct {
	ID         primitive.ObjectID   `bson:"_id" json:"id"`
	Orders     []primitive.ObjectID `bson:"orders" json:"orders"`
	RouteOrder []int                `bson:"routeOrder" json:"routeOrder"`
	Number     string               `bson:"number" json:"number"`
	Finalized  bool                 `bson:"finalized" json:"finalized"`
	Delivered  bool                 `bson:"delivered" json:"delivered"`
	Driver     primitive.ObjectID   `bson:"driver" json:"driver"`
	DriverName string               `bson:"driverName" json:"driverName"`
	Store      primitive.ObjectID   `bson:"store" json:"store"`
	CreatedAt  time.Time            `bson:"createdAt" json:"createdAt"`
	UpdatedAt  time.Time            `bson:"updatedAt" json:"updatedAt"`
}
