package product

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/asia-loop-gmbh/lambda-utils-go/v4/pkg/servicewoo"
)

func GetVariation(ctx context.Context, productID int) ([]servicewoo.ProductVariation, error) {
	woo, err := servicewoo.NewWoo(ctx)
	if err != nil {
		return nil, err
	}
	url := woo.NewURL(ctx, fmt.Sprintf("/products/%d/variations", productID))
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	variations := make([]servicewoo.ProductVariation, 0)
	err = json.Unmarshal(body, &variations)
	if err != nil {
		return nil, err
	}
	return variations, nil
}
