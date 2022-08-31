package servicesns_test

import (
	"context"
	"testing"

	commoncontext "github.com/nam-truong-le/lambda-utils-go/pkg/context"
	"github.com/stretchr/testify/assert"

	"github.com/asia-loop-gmbh/lambda-utils-go/v4/pkg/servicesns"
)

func TestPublishOrderCreated(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	ctx := context.WithValue(context.TODO(), commoncontext.FieldStage, "dev")
	err := servicesns.PublishOrderCreated(ctx, &servicesns.EventOrderCreatedData{
		OrderID: "POS-810052",
	})
	assert.NoError(t, err)
}
