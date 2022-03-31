package product

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicewoo"
)

const (
	perPage = 100
)

func Get(log *logrus.Entry, ctx context.Context, stage string) ([]servicewoo.Product, error) {
	log.Infof("get all products from woo")
	page := 1
	result := make([]servicewoo.Product, 0)
	for true {
		ps, err := getPage(log, ctx, stage, page)
		if err != nil {
			return nil, err
		}
		if len(ps) == 0 {
			break
		}
		for _, p := range ps {
			if p.Status == "publish" {
				result = append(result, p)
			}
		}
		page++
	}
	return result, nil
}

func getPage(log *logrus.Entry, ctx context.Context, stage string, page int) ([]servicewoo.Product, error) {
	log.Infof("get products from woo page [%d], per_page [%d]", page, perPage)
	woo, err := servicewoo.NewWoo(log, ctx, stage)
	if err != nil {
		return nil, err
	}

	url := woo.NewURL(log, fmt.Sprintf("/products?page=%d&per_page=%d", page, perPage))
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("woo returns [%s]", res.Status)
	}
	products := make([]servicewoo.Product, 0)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &products)
	return products, nil
}
