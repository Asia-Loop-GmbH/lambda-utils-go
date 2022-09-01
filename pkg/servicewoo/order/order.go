package order

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/asia-loop-gmbh/lambda-utils-go/v4/pkg/servicewoo"
	"github.com/nam-truong-le/lambda-utils-go/pkg/logger"
)

func GetRefunds(ctx context.Context, id int) ([]servicewoo.Refund, error) {
	log := logger.FromContext(ctx)
	w, err := servicewoo.NewWoo(ctx)
	if err != nil {
		log.Errorf("%v", err)
		return nil, err
	}
	url := w.NewURL(ctx, fmt.Sprintf("/orders/%d/refunds", id))
	res, err := http.Get(url)
	if err != nil {
		log.Errorf("%v", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Warnf("failed to close response body: %s", err)
		}
	}(res.Body)
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

func Get(ctx context.Context, id int) (*servicewoo.Order, error) {
	log := logger.FromContext(ctx)
	w, err := servicewoo.NewWoo(ctx)
	if err != nil {
		log.Errorf("%v", err)
		return nil, err
	}
	url := w.NewURL(ctx, fmt.Sprintf("/orders/%d", id))
	res, err := http.Get(url)
	if err != nil {
		log.Errorf("%v", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Warnf("failed to close response body: %s", err)
		}
	}(res.Body)
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
