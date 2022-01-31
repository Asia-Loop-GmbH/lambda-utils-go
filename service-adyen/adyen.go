package service_adyen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/adyen/adyen-go-api-library/v5/src/adyen"
	"github.com/adyen/adyen-go-api-library/v5/src/common"
	"github.com/asia-loop-gmbh/lambda-utils-go/aws"
	"github.com/asia-loop-gmbh/lambda-utils-go/text"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	envMap = map[string]common.Environment{
		"dev":  common.TestEnv,
		"pre":  common.LiveEnv,
		"prod": common.LiveEnv,
	}
)

func newClient(stage string) (*adyen.APIClient, error) {
	apiKey, err := aws.GetSSMParameter(stage, "/adyen/key", true)
	if err != nil {
		return nil, err
	}
	environment, ok := envMap[stage]
	if !ok {
		return nil, fmt.Errorf("no adyen environment config found for stage: %s", stage)
	}
	client := adyen.NewClient(&common.Config{
		ApiKey:      *apiKey,
		Environment: environment,
	})
	return client, nil
}

func NewTender(stage, pos, orderId string, amount float32) error {
	client, err := newClient(stage)
	if err != nil {
		return err
	}

	terminalRequest := TerminalAPIRequest{
		SaleToPOIRequest: SaleToPOIRequest{
			MessageHeader: MessageHeader{
				ProtocolVersion: protocolVersion,
				MessageClass:    messageClassService,
				MessageCategory: messageCategoryPayment,
				MessageType:     messageTypeRequest,
				SaleID:          orderId,
				ServiceID:       text.RandomString(10, false, false, true),
				POIID:           pos,
			},
			PaymentRequest: PaymentRequest{
				SaleData: SaleData{
					SaleTransactionID: SaleTransactionID{
						TransactionID: orderId,
						TimeStamp:     time.Now().Format("2006-01-02T15:04:05-07:00"),
					},
				},
				PaymentTransaction: PaymentTransaction{
					AmountsReq: AmountsReq{
						Currency:        currencyEUR,
						RequestedAmount: amount,
					},
				},
			},
		},
	}

	requestBody, err := json.Marshal(terminalRequest)
	if err != nil {
		return err
	}
	url := fmt.Sprintf("%s/%s", client.GetConfig().TerminalApiCloudEndpoint, "async")
	log.Printf("create tender: POST %s", url)
	httpClient := http.Client{}
	postRequest, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	postRequest.Header.Set("Content-Type", "application/json")
	postRequest.Header.Set("x-API-key", client.GetConfig().ApiKey)
	response, err := httpClient.Do(postRequest)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode >= 300 {
		responseBody, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("adyen request error %s: %s", response.Status, string(responseBody))
	}

	log.Printf("status %s", response.Status)

	return nil
}