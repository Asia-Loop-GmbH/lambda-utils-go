package service_slots

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/asia-loop-gmbh/lambda-types-go/woo"
	servicewoo "github.com/asia-loop-gmbh/lambda-utils-go/v3/service-woo"
)

func GetSlots(log *logrus.Entry, stage string) (*woo.Slots, error) {
	log.Info("get slots")
	serviceWoo, err := servicewoo.NewWoo(log, stage)
	if err != nil {
		return nil, err
	}
	response, err := http.Get(serviceWoo.NewURLAsiaLoop(log, "/delivery-slots"))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	slots := new(woo.Slots)
	if err := json.Unmarshal(responseBody, slots); err != nil {
		return nil, err
	}

	return slots, nil
}
