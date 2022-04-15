package service

import (
	"time"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/dbadmin"
)

const (
	WooTaxClass7 = "mitnehmen"
	WooFeeTip    = "Tip"
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
	CreatedAt      string                      `dynamodbav:"CreatedAt"`
	Type           RevenueType                 `dynamodbav:"RevenueType"`
	ShippingMethod dbadmin.OrderShippingMethod `dynamodbav:"ShippingMethod"`
	Store          string                      `dynamodbav:"Store"`
	Source         RevenueSource               `dynamodbav:"Source"`
	Company        string                      `dynamodbav:"Company"`
	Net7           string                      `dynamodbav:"Net7"`
	Tax7           string                      `dynamodbav:"Tax7"`
	Tip            string                      `dynamodbav:"Tip"`
}

func TimeToDynamoString(t time.Time) (*string, error) {
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		return nil, err
	}
	result := t.In(loc).Format("2006-01-02T15:04:05Z07")
	return &result, nil
}
