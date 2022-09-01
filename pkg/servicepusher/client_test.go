package servicepusher

import (
	"context"
	"testing"

	commoncontext "github.com/nam-truong-le/lambda-utils-go/pkg/context"
	"github.com/stretchr/testify/assert"
)

func TestGetClient(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	ctx := context.WithValue(context.Background(), commoncontext.FieldStage, "dev")

	c, err := getClient(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, c)

	c, err = getClient(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, c)
}
