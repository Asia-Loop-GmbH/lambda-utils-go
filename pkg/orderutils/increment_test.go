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
	ctx := context.WithValue(context.Background(), commoncontext.FieldStage, "dev")
	next, err := orderutils.NextOrderInvoice(ctx)
	assert.NoError(t, err)
	log.Printf("%s", *next)
}

func TestNextOrderInvoiceLieferando(t *testing.T) {
	ctx := context.WithValue(context.Background(), commoncontext.FieldStage, "dev")
	next, err := orderutils.NextOrderInvoiceLieferando(ctx)
	assert.NoError(t, err)
	log.Printf("%s", *next)
}
