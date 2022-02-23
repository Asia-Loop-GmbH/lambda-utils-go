package servicewoo

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"

	"strings"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicessm"
)

type Woo struct {
	URL    string
	Key    string
	Secret string
}

func NewWoo(log *logrus.Entry, ctx context.Context, stage string) (*Woo, error) {
	log.Infof("read woo information")
	shopUrl, err := servicessm.GetParameter(log, ctx, stage, "/shop/url", false)
	if err != nil {
		return nil, err
	}
	wooKey, err := servicessm.GetParameter(log, ctx, stage, "/shop/woo/key", false)
	if err != nil {
		return nil, err
	}
	wooSecret, err := servicessm.GetParameter(log, ctx, stage, "/shop/woo/secret", true)
	if err != nil {
		return nil, err
	}

	return &Woo{*shopUrl, *wooKey, *wooSecret}, nil
}

func (w *Woo) NewURL(log *logrus.Entry, url string) string {
	return w.newURL(log, url, "/wp-json/wc/v3")
}

func (w *Woo) NewURLAsiaLoop(log *logrus.Entry, url string) string {
	return w.newURL(log, url, "/wp-json/asialoop-api")
}

func (w *Woo) newURL(log *logrus.Entry, url string, api string) string {
	log.Infof("prepare woo url: %s", url)
	connector := "?"
	if strings.Contains(url, "?") {
		connector = "&"
	}
	result := fmt.Sprintf(
		"%s%s%s%sconsumer_key=%s&consumer_secret=%s",
		w.URL, api, url, connector, w.Key, w.Secret,
	)
	log.Infof("final woo url: %s", result)
	return result
}
