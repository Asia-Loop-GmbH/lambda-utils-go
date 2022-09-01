package serviceadyen_test

import (
	"context"
	"testing"

	commoncontext "github.com/nam-truong-le/lambda-utils-go/pkg/context"
	"github.com/stretchr/testify/assert"

	"github.com/asia-loop-gmbh/lambda-utils-go/v4/pkg/random"
	"github.com/asia-loop-gmbh/lambda-utils-go/v4/pkg/serviceadyen"
)

func TestNewTender(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	orderId := random.String(6, true, true, true)
	ctx := context.WithValue(context.TODO(), commoncontext.FieldStage, "dev")
	err := serviceadyen.NewTender(ctx, "S1F2-000158213300585", orderId, 10.12)
	assert.Nil(t, err)
}
