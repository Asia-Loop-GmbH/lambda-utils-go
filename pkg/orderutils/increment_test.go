package orderutils_test

import (
	"context"
	"log"
	"testing"

	commoncontext "github.com/nam-truong-le/lambda-utils-go/pkg/context"
	"github.com/stretchr/testify/assert"

	"github.com/asia-loop-gmbh/lambda-utils-go/v4/pkg/orderutils"
)

func TestNextOrderInvoice(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	ctx := context.WithValue(context.Background(), commoncontext.FieldStage, "dev")
	next, err := orderutils.NextOrderInvoice(ctx)
	assert.NoError(t, err)
	log.Printf("%s", *next)
}
