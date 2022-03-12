package revenue

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/revenue/service"
)

const (
	nonWooRefundSuffix = "--non-woo-refund"
)

func fakeNonWooRefund(log *logrus.Entry, ctx context.Context, stage, paymentID string) error {
	rs, err := service.QueryByPaymentID(log, ctx, stage, paymentID)
	if err != nil {
		return err
	}

	if len(rs) == 0 {
		// TODO: currently not support corporate refund
		// so it must be POS refund here
		if err := SyncPOSOrderByUUID(log, ctx, stage, paymentID); err != nil {
			return err
		}
		rs, err = service.QueryByPaymentID(log, ctx, stage, paymentID)
		if err != nil {
			return err
		}
	}

	if len(rs) != 1 {
		return fmt.Errorf("expected only one revenue item from [%s], got: %v", paymentID, rs)
	}

	r := rs[0]

	if r.Type != service.RevenueTypeOrder {
		return fmt.Errorf("expected only order from ref [%s], got: %v", paymentID, r)
	}

	refund := service.Revenue{
		ID:             r.ID + nonWooRefundSuffix,
		PaymentID:      r.PaymentID,
		CreatedAt:      r.CreatedAt, // TODO: get date form adyen
		Type:           service.RevenueTypeRefund,
		ShippingMethod: r.ShippingMethod,
		Store:          r.Store,
		Source:         r.Source,
		Company:        r.Company,
		Net7:           "-" + r.Net7,
		Tax7:           "-" + r.Tax7,
		Tip:            "-" + r.Tip,
	}

	return insert(log, ctx, stage, &refund)
}
