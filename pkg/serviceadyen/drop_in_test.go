package serviceadyen_test

import (
	"context"
	"testing"

	context2 "github.com/nam-truong-le/lambda-utils-go/pkg/context"
	"github.com/stretchr/testify/assert"

	"github.com/asia-loop-gmbh/lambda-utils-go/v4/pkg/random"
	"github.com/asia-loop-gmbh/lambda-utils-go/v4/pkg/serviceadyen"
)

func TestNewDropInPayment_Success(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	ctx := context.WithValue(context.TODO(), context2.FieldStage, "dev")
	response, err := serviceadyen.NewDropInPayment(
		ctx,
		"10.23",
		random.String(10, true, true, true),
		"https://admin2-dev.asia-loop.com",
	)
	assert.NoError(t, err)
	assert.NotNil(t, response)
}
