package coupon

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"

	"github.com/asia-loop-gmbh/lambda-types-go/v2/pkg/woo"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicewoo"

	"io/ioutil"
	"net/http"
	"strings"
)

func IsValidAndHasEnough(log *logrus.Entry, ctx context.Context, stage, code, appliedAmount string) bool {
	log.Infof("check coupon valid and has enough amount: %s", code)
	coupon, err := GetCouponByCode(log, ctx, stage, code)
	if err != nil {
		log.Errorf("could not get coupon '%s': %s", code, err)
		return false
	}
	current, err := decimal.NewFromString(coupon.Amount)
	if err != nil {
		return false
	}
	toUse, err := decimal.NewFromString(appliedAmount)
	if err != nil {
		return false
	}
	return current.Cmp(toUse) > 0
}

func GetCouponByCode(log *logrus.Entry, ctx context.Context, stage, code string) (*woo.Coupon, error) {
	log.Infof("get coupon: %s", code)
	code = strings.TrimSpace(code)
	if code == "" {
		return nil, fmt.Errorf("blank coupon code")
	}
	serviceWoo, err := servicewoo.NewWoo(log, ctx, stage)
	if err != nil {
		return nil, err
	}

	response, err := http.Get(serviceWoo.NewURL(log, fmt.Sprintf("/coupons?code=%s", code)))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	coupons := make([]woo.Coupon, 0)
	if err := json.Unmarshal(responseBody, &coupons); err != nil {
		return nil, err
	}

	if len(coupons) == 0 {
		return nil, fmt.Errorf("coupon code '%s' not found", code)
	}
	if len(coupons) > 1 {
		return nil, fmt.Errorf("multiple coupon codes found for '%s'", code)
	}
	return &coupons[0], nil
}

func UpdateCouponByCode(log *logrus.Entry, ctx context.Context, stage, code, amount string) error {
	log.Infof("update coupon %s: %s", code, amount)
	coupon, err := GetCouponByCode(log, ctx, stage, code)
	if err != nil {
		return err
	}
	toUse, err := decimal.NewFromString(amount)
	if err != nil {
		return err
	}
	currentAmount, err := decimal.NewFromString(coupon.Amount)
	if err != nil {
		return err
	}
	newAmount := currentAmount.Sub(toUse)

	updateCoupon := woo.Coupon{
		Amount: newAmount.StringFixed(2),
	}

	serviceWoo, err := servicewoo.NewWoo(log, ctx, stage)
	if err != nil {
		return err
	}

	requestBody, err := json.Marshal(updateCoupon)
	if err != nil {
		return err
	}
	log.Infof("update coupon request body: %s", string(requestBody))

	url := serviceWoo.NewURL(log, fmt.Sprintf("/coupons/%d", coupon.ID))
	log.Infof("PUT -> %s", url)
	httpClient := http.Client{}
	request, err := http.NewRequest(
		http.MethodPut,
		url,
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := httpClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	log.Infof("response: %s", string(responseBody))

	if response.StatusCode >= 300 {
		return fmt.Errorf("could not update coupon, error '%d': %s", response.StatusCode, string(responseBody))
	}

	return nil
}
