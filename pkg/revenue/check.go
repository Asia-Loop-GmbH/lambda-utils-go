package revenue

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/revenue/service"
)

func refundExists(log *logrus.Entry, ctx context.Context, stage string, merchantRef string) (bool, error) {
	rs, err := service.QueryByPaymentID(log, ctx, stage, merchantRef)
	if err != nil {
		return false, err
	}

	if len(rs) == 0 {
		return false, nil
	}

	for _, r := range rs {
		if r.Type == service.RevenueTypeRefund {
			return true, nil
		}
	}

	return false, nil
}

func orderExists(log *logrus.Entry, ctx context.Context, stage string, merchantRef string) (bool, error) {
	_, err := service.GetByID(log, ctx, stage, merchantRef)
	if err == service.ErrorNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
