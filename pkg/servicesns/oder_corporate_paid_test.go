package servicesns_test

import (
	"context"
	"testing"

	commoncontext "github.com/nam-truong-le/lambda-utils-go/pkg/context"
	"github.com/stretchr/testify/assert"

	"github.com/asia-loop-gmbh/lambda-utils-go/v4/pkg/servicesns"
)

func TestPublishOrderCorporatePaid(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	ctx := context.WithValue(context.TODO(), commoncontext.FieldStage, "dev")
	err := servicesns.PublishOrderCorporatePaid(ctx, &servicesns.EventOrderCorporatePaidData{
		OrderID: "POS-9C563XKE",
	})
	assert.NoError(t, err)
}
