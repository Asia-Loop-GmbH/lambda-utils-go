package api

import "github.com/asia-loop-gmbh/lambda-utils-go/v4/pkg/dbadmin"

type OrderDetails struct {
	Order    dbadmin.Order    `json:"order"`
	Customer dbadmin.Customer `json:"customer"`
}

type SearchOrderRequest struct {
	Text  *string `json:"text"`
	Limit *int64  `json:"limit"`
}
