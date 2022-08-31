package servicemongo_test

import (
	"context"
	"testing"

	commoncontext "github.com/nam-truong-le/lambda-utils-go/pkg/context"
	"github.com/stretchr/testify/assert"

	"github.com/asia-loop-gmbh/lambda-utils-go/v4/pkg/servicemongo"
)

func TestNextByStage(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	ctx := context.WithValue(context.Background(), commoncontext.FieldStage, "dev")
	next, err := servicemongo.Next(ctx, "test")
	assert.NoError(t, err)
	assert.True(t, next > 0)
	servicemongo.Disconnect(ctx)
}
