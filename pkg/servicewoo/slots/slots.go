package slots

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/asia-loop-gmbh/lambda-utils-go/v4/pkg/servicewoo"
	"github.com/nam-truong-le/lambda-utils-go/pkg/logger"
)

func GetSlots(ctx context.Context) (*Slots, error) {
	log := logger.FromContext(ctx)
	log.Info("get slots")
	serviceWoo, err := servicewoo.NewWoo(ctx)
	if err != nil {
		return nil, err
	}
	response, err := http.Get(serviceWoo.NewURLAsiaLoop(ctx, "/delivery-slots"))
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Warnf("failed to close response body: %s", err)
		}
	}(response.Body)
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	slots := new(Slots)
	if err := json.Unmarshal(responseBody, slots); err != nil {
		return nil, err
	}

	return slots, nil
}
