package api

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/dbadmin"
)

type OrderPublic struct {
	ID            primitive.ObjectID  `json:"id"`
	OrderID       string              `json:"orderId"`
	OrderNumber   string              `json:"orderNumber"`
	InvoiceNumber string              `json:"invoiceNumber"`
	Address       string              `json:"address"`
	AddressLine2  string              `json:"addressLine2"`
	DeliveryDate  string              `json:"deliveryDate"`
	DeliveryTime  string              `json:"deliveryTime"`
	Status        dbadmin.OrderStatus `json:"status"`
	Printed       bool                `json:"printed"`
	CreatedAt     time.Time           `json:"createdAt"`
	GroupKey      string              `json:"groupKey"`
	GroupPosition int                 `json:"groupPosition"`
	DriverName    string              `json:"driverName"`
	Invoice       string              `json:"invoice"`
}
