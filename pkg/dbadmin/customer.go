package dbadmin

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	CustomerRef  string             `bson:"customerRef" json:"customerRef"`
	FirstName    string             `bson:"firstName" json:"firstName"`
	LastName     string             `bson:"lastName" json:"lastName"`
	AddressLine1 string             `bson:"addressLine1" json:"addressLine1"`
	AddressLine2 string             `bson:"addressLine2" json:"addressLine2"`
	Postcode     string             `bson:"postcode" json:"postcode"`
	City         string             `bson:"city" json:"city"`
	Telephone    string             `bson:"telephone" json:"telephone"`
	Email        string             `bson:"email" json:"email"`
	Boxes        []int              `bson:"boxes" json:"boxes"`
	CreatedAt    time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt" json:"updatedAt"`
}
