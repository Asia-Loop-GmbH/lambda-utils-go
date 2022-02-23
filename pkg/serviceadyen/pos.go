package serviceadyen

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/random"
)

func NewTender(log *logrus.Entry, ctx context.Context, stage, pos, orderId string, amount float32) error {
	log.Infof("new POS payment in %s: order %s [%f]", stage, orderId, amount)
	client, err := newClient(log, ctx, stage)
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
				ServiceID:       random.String(10, false, false, true),
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
	log.Infof("create tender: POST %s", url)
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

	log.Infof("status %s", response.Status)

	return nil
}
