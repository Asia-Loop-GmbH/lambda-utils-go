package api

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/dbadmin"
)

type CreateOrderRequest struct {
	FirstName          string                      `json:"firstName"`
	LastName           string                      `json:"lastName"`
	Telephone          string                      `json:"telephone"`
	Email              string                      `json:"email"`
	AddressLine1       string                      `json:"addressLine1"`
	AddressLine2       string                      `json:"addressLine2"`
	Postcode           string                      `json:"postcode"`
	City               string                      `json:"city"`
	ShippingMethod     dbadmin.OrderShippingMethod `json:"shippingMethod"`
	Items              []dbadmin.OrderItem         `json:"items"`
	Net                string                      `json:"net"`
	Tax                string                      `json:"tax"`
	Total              string                      `json:"total"`
	CouponCode         string                      `json:"couponCode"`
	AppliedCouponNet   string                      `json:"appliedCouponNet"`
	AppliedCouponTax   string                      `json:"appliedCouponTax"`
	AppliedCouponTotal string                      `json:"appliedCouponTotal"`
	Store              primitive.ObjectID          `json:"store"`
	CompanyKey         string                      `json:"companyKey"`
}

type CreateOrderAddressOptions struct {
	FirstName    string
	LastName     string
	Telephone    string
	Email        string
	AddressLine1 string
	AddressLine2 string
	Postcode     string
	City         string
}

type CreateOrderOrderOptions struct {
	OrderID            string
	OrderNumber        string
	InvoiceNumber      string
	Items              []dbadmin.OrderItem
	DeliveryDate       string
	DeliveryTime       string
	Net                string
	Total              string
	Tax                string
	Tip                string
	CouponCode         string
	AppliedCouponNet   string
	AppliedCouponTax   string
	AppliedCouponTotal string
	ShippingMethod     dbadmin.OrderShippingMethod
	CustomerNote       string
	Status             dbadmin.OrderStatus
	Store              primitive.ObjectID
	CompanyKey         string
}
