package dbadmin

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TicketCustomer struct {
	FirstName      string `bson:"firstName" json:"firstName"`
	LastName       string `bson:"lastName" json:"lastName"`
	Telephone      string `bson:"telephone" json:"telephone"`
	Email          string `bson:"email" json:"email"`
	OrderReference string `bson:"orderReference" json:"orderReference"`
}

type TicketMessage struct {
	From      string    `bson:"from" json:"from"`
	Role      string    `bson:"role" json:"role"`
	Text      string    `bson:"text" json:"text"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
}

type Ticket struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	TicketNumber string             `bson:"ticketNumber" json:"ticketNumber"`
	StoreKey     string             `bson:"storeKey" json:"storeKey"`
	Customer     TicketCustomer     `bson:"customer" json:"customer"`
	CreatedAt    time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt" json:"updatedAt"`
	Messages     []TicketMessage    `bson:"messages" json:"messages"`
}
