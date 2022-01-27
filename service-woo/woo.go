package service_woo

import (
	"fmt"
	utils "github.com/asia-loop-gmbh/lambda-utils-go"
	"strings"
)

type Woo struct {
	URL    string
	Key    string
	Secret string
}

func NewWoo(stage string) (*Woo, error) {
	shopUrl, err := utils.GetSSMParameter(stage, "/shop/url", false)
	if err != nil {
		return nil, err
	}
	wooKey, err := utils.GetSSMParameter(stage, "/shop/woo/key", false)
	if err != nil {
		return nil, err
	}
	wooSecret, err := utils.GetSSMParameter(stage, "/shop/woo/secret", true)
	if err != nil {
		return nil, err
	}

	return &Woo{*shopUrl, *wooKey, *wooSecret}, nil
}

func (w *Woo) NewURL(url string) string {
	connector := "?"
	if strings.Contains(url, "?") {
		connector = "&"
	}
	return fmt.Sprintf(
		"%s/wp-json/wc/v3%s%sconsumer_key=%s&consumer_secret=%s",
		w.URL, url, connector, w.Key, w.Secret,
	)
}
