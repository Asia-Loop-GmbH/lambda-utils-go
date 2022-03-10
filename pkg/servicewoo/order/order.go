package order

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicewoo"
)

func GetRefunds(log *logrus.Entry, ctx context.Context, stage string, id int) ([]servicewoo.Refund, error) {
	w, err := servicewoo.NewWoo(log, ctx, stage)
	if err != nil {
		log.Errorf("%v", err)
		return nil, err
	}
	url := w.NewURL(log, fmt.Sprintf("/orders/%d/refunds", id))
	res, err := http.Get(url)
	if err != nil {
		log.Errorf("%v", err)
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Errorf("status %d", res.StatusCode)
		return nil, fmt.Errorf("status %d", res.StatusCode)
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf("%v", err)
		return nil, err
	}
	refunds := make([]servicewoo.Refund, 0)
	err = json.Unmarshal(resBody, &refunds)
	if err != nil {
		log.Errorf("%v", err)
		return nil, err
	}
	return refunds, nil
}

func Get(log *logrus.Entry, ctx context.Context, stage string, id int) (*servicewoo.Order, error) {
	w, err := servicewoo.NewWoo(log, ctx, stage)
	if err != nil {
		log.Errorf("%v", err)
		return nil, err
	}
	url := w.NewURL(log, fmt.Sprintf("/orders/%d", id))
	res, err := http.Get(url)
	if err != nil {
		log.Errorf("%v", err)
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Errorf("status %d", res.StatusCode)
		return nil, fmt.Errorf("status %d", res.StatusCode)
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf("%v", err)
		return nil, err
	}
	wooOrder := new(servicewoo.Order)
	err = json.Unmarshal(resBody, wooOrder)
	if err != nil {
		log.Errorf("%v", err)
		return nil, err
	}
	return wooOrder, nil
}
