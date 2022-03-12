package revenue

import (
	"context"
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/dbadmin"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/revenue/service"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicewoo"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicewoo/order"
)

func SyncWooRefunds(log *logrus.Entry, ctx context.Context, stage string, id int) error {
	o, err := order.Get(log, ctx, stage, id)
	if err != nil {
		return err
	}

	if len(o.Refunds) == 0 {
		log.Warnf("there is no refund to process")
		return nil
	}

	refunds, err := order.GetRefunds(log, ctx, stage, id)
	if err != nil {
		return err
	}

	for _, re := range refunds {
		if err := syncWooRefund(log, ctx, stage, o, &re); err != nil {
			log.Errorf("%s", err)
			return err
		}
		log.Infof("refund [%d] from order [%d] processed", re.ID, id)
	}

	return nil
}

func syncWooRefund(log *logrus.Entry, ctx context.Context, stage string, o *servicewoo.Order, re *servicewoo.Refund) error {
	net7 := decimal.Zero
	tax7 := decimal.Zero

	for _, item := range re.LineItems {
		switch item.TaxClass {
		case service.WooTaxClass7:
			net, err := decimal.NewFromString(item.Total)
			if err != nil {
				return err
			}
			tax, err := decimal.NewFromString(item.TotalTax)
			if err != nil {
				return err
			}
			net7 = net7.Add(net)
			tax7 = tax7.Add(tax)
		default:
			return fmt.Errorf("tax class [%s] not registered", item.TaxClass)
		}
	}

	tip := decimal.Zero
	for _, item := range re.FeeLines {
		if strings.HasPrefix(item.Name, service.WooFeeTip) {
			tipTax, err := decimal.NewFromString(item.TotalTax)
			if err != nil {
				return err
			}
			if !tipTax.IsZero() {
				return fmt.Errorf("tip tax from refund [%d] must be zero, got [%s]", re.ID, tipTax.StringFixed(2))
			}
			tipTotal, err := decimal.NewFromString(item.Total)
			if err != nil {
				return err
			}
			tip = tip.Add(tipTotal)
		} else {
			return fmt.Errorf("fee name [%s] not registered", item.Name)
		}
	}

	total, err := decimal.NewFromString(re.Amount)
	if err != nil {
		return err
	}
	checkTotal := net7.Add(tax7).Add(tip)
	if !total.Neg().Equal(checkTotal) {
		return fmt.Errorf("expected total [%s] from refund [%d], calculated [%s]", total.StringFixed(2), re.ID, checkTotal.StringFixed(2))
	}

	created, err := re.GetDateCreated()
	if err != nil {
		return err
	}

	createdString, err := service.TimeToDynamoString(created)
	if err != nil {
		return err
	}
	r := service.Revenue{
		ID:             fmt.Sprintf("%d", re.ID),
		PaymentID:      fmt.Sprintf("%d", o.ID),
		CreatedAt:      *createdString,
		ShippingMethod: dbadmin.OrderShippingMethod(o.GetShippingMethod()),
		Store:          o.GetStoreKey(),
		Source:         service.RevenueSourceOnline,
		Company:        "",
		Type:           service.RevenueTypeRefund,
		Net7:           net7.StringFixed(2),
		Tax7:           tax7.StringFixed(2),
		Tip:            tip.StringFixed(2),
	}

	return insert(log, ctx, stage, &r)
}
