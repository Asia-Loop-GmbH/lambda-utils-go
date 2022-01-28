package service_coupon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/asia-loop-gmbh/lambda-types-go/woo"
	utils "github.com/asia-loop-gmbh/lambda-utils-go/number"
	servicewoo "github.com/asia-loop-gmbh/lambda-utils-go/service-woo"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"strings"
)

func IsValidAndHasEnough(stage, code, appliedAmount string) bool {
	coupon, err := GetCouponByCode(stage, code)
	if err != nil {
		log.Printf("could not get coupon '%s': %s", code, err)
		return false
	}
	current, err := utils.NewBigNumber(coupon.Amount)
	if err != nil {
		return false
	}
	toUse, err := utils.NewBigNumber(appliedAmount)
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
	toUse, err := utils.NewBigNumber(amount)
	if err != nil {
		return err
	}
	currentAmount, err := utils.NewBigNumber(coupon.Amount)
	if err != nil {
		return err
	}
	newAmount := &big.Float{}
	newAmount.Sub(currentAmount, toUse)

	coupon.Amount = newAmount.Text('f', 2)

	serviceWoo, err := servicewoo.NewWoo(stage)
	if err != nil {
		return err
	}

	requestBody, err := json.Marshal(coupon)
	if err != nil {
		return err
	}

	httpClient := http.Client{}
	request, err := http.NewRequest(
		"PUT",
		serviceWoo.NewURL(fmt.Sprintf("/coupons/%d", coupon.ID)),
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return err
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 300 {
		return fmt.Errorf("could not update coupon, error '%d': %s", response.StatusCode, string(responseBody))
	}

	return nil
}
