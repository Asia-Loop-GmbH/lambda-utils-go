package service_coupon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/asia-loop-gmbh/lambda-types-go/woo"
	servicewoo "github.com/asia-loop-gmbh/lambda-utils-go/service-woo"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func IsValidAndHasEnough(stage, code, appliedAmount string) bool {
	coupon, err := GetCouponByCode(stage, code)
	if err != nil {
		log.Printf("could not get coupon '%s': %s", code, err)
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

func GetCouponByCode(stage, code string) (*woo.Coupon, error) {
	code = strings.TrimSpace(code)
	if code == "" {
		return nil, fmt.Errorf("blank coupon code")
	}
	serviceWoo, err := servicewoo.NewWoo(stage)
	if err != nil {
		return nil, err
	}

	response, err := http.Get(serviceWoo.NewURL(fmt.Sprintf("/coupons?code=%s", code)))
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

func UpdateCouponByCode(stage, code, amount string) error {
	coupon, err := GetCouponByCode(stage, code)
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

	serviceWoo, err := servicewoo.NewWoo(stage)
	if err != nil {
		return err
	}

	requestBody, err := json.Marshal(updateCoupon)
	if err != nil {
		return err
	}
	log.Printf("update coupon request body: %s", string(requestBody))

	url := serviceWoo.NewURL(fmt.Sprintf("/coupons/%d", coupon.ID))
	log.Printf("PUT -> %s", url)
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
	log.Printf("response: %s", string(responseBody))

	if response.StatusCode >= 300 {
		return fmt.Errorf("could not update coupon, error '%d': %s", response.StatusCode, string(responseBody))
	}

	return nil
}
