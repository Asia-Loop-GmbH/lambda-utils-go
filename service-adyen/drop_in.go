package service_adyen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/adyen/adyen-go-api-library/v5/src/checkout"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"

	"github.com/asia-loop-gmbh/lambda-types-go/admin"
)

type SessionResponse struct {
	ID          string `json:"id"`
	SessionData string `json:"sessionData"`
}

func NewDropInPayment(log *logrus.Entry, stage, value, ref, returnURL string) (*admin.PaymentDropInResponse, error) {
	log.Infof("new drop in payment for order [%s]", ref)
	client, err := newClient(log, stage)
	if err != nil {
		return nil, err
	}

	amount, err := decimal.NewFromString(value) // TODO: only valid for corporate customers!!!
	if err != nil {
		return nil, err
	}
	amountInt := amount.Mul(decimal.NewFromInt(100)).IntPart()

	req := &checkout.PaymentSetupRequest{
		Amount:          checkout.Amount{Currency: currencyEUR, Value: amountInt},
		MerchantAccount: accountECOM,
		ReturnUrl:       returnURL,
		Reference:       ref,
		CountryCode:     countryDE,
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/v68/sessions", client.GetConfig().CheckoutEndpoint)
	log.Printf("POST -> %s", url)
	httpClient := http.Client{}
	postRequest, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	postRequest.Header.Set("x-API-key", client.GetConfig().ApiKey)
	response, err := httpClient.Do(postRequest)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode >= 300 {
		responseBody, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("adyen request error %s: %s", response.Status, string(responseBody))
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	newSession := new(SessionResponse)
	if err := json.Unmarshal(responseBody, newSession); err != nil {
		return nil, err
	}
	result := admin.PaymentDropInResponse{
		ID:          newSession.ID,
		SessionData: newSession.SessionData,
	}
	log.Infof("%v", result)
	return &result, nil
}
