package servicesns_test

import (
	"context"
	"testing"

	commoncontext "github.com/nam-truong-le/lambda-utils-go/pkg/context"
	"github.com/stretchr/testify/assert"

	"github.com/asia-loop-gmbh/lambda-utils-go/v4/pkg/servicesns"
)

func TestPublishOrderPickupReady(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	ctx := context.WithValue(context.Background(), commoncontext.FieldStage, "dev")
	err := servicesns.PublishOrderPickupReady(ctx, &servicesns.EventOrderPickupReadyData{
		OrderID: "POS-810052",
		InTime:  "10 Minuten",
	})
	assert.NoError(t, err)
}
