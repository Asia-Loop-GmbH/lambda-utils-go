package revenue

import (
	"time"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/dbadmin"
)

const (
	WooTaxClass7 = "mitnehmen"
	WooFeeTip    = "Tip (Benutzerdefiniertes Trinkgeld)"
)

type RevenueSource string

const (
	RevenueSourceOnline    RevenueSource = "Online"
	RevenueSourceOffline   RevenueSource = "POS"
	RevenueSourceCorporate RevenueSource = "Corporate"
)

type RevenueType string

const (
	RevenueTypeOrder  RevenueType = "Order"
	RevenueTypeRefund RevenueType = "Refund"
)

type Revenue struct {
	ID             string                      `dynamodbav:"Id"`
	PaymentID      string                      `dynamodbav:"PaymentId"`
	CreatedAt      time.Time                   `dynamodbav:"CreatedAt"`
	Type           RevenueType                 `dynamodbav:"RevenueType"`
	ShippingMethod dbadmin.OrderShippingMethod `dynamodbav:"ShippingMethod"`
	Store          string                      `dynamodbav:"Store"`
	Source         RevenueSource               `dynamodbav:"Source"`
	Company        string                      `dynamodbav:"Company"`
	Net7           string                      `dynamodbav:"Net7"`
	Tax7           string                      `dynamodbav:"Tax7"`
	Tip            string                      `dynamodbav:"Tip"`
}
