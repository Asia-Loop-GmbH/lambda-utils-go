package orderutils

import (
	"context"
	"fmt"
	"time"

	"github.com/asia-loop-gmbh/lambda-utils-go/v4/pkg/servicemongo"
)

const (
	incrementOrderInvoiceKey = "ORDER_INVOICE"
)

func NextOrderInvoice(ctx context.Context) (*string, error) {
	next, err := servicemongo.Next(ctx, incrementOrderInvoiceKey)
	if err != nil {
		return nil, err
	}
	prefix, err := orderInvoicePrefix()
	if err != nil {
		return nil, err
	}
	full := fmt.Sprintf("P%s-%07d", *prefix, next)
	return &full, nil
}

func orderInvoicePrefix() (*string, error) {
	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		return nil, err
	}
	prefix := time.Now().In(loc).Format("200601")
	return &prefix, nil
}
