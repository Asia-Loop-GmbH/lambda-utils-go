package servicemongo_test

import (
	"context"
	"testing"

	commoncontext "github.com/nam-truong-le/lambda-utils-go/pkg/context"
	"github.com/stretchr/testify/assert"

	"github.com/asia-loop-gmbh/lambda-utils-go/v4/pkg/dbadmin"
	"github.com/asia-loop-gmbh/lambda-utils-go/v4/pkg/servicemongo"
)

func TestAdminCollection(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	ctx := context.WithValue(context.Background(), commoncontext.FieldStage, "dev")
	_, err := servicemongo.AdminCollection(ctx, dbadmin.CollectionOrder)
	assert.NoError(t, err)
	_, err = servicemongo.AdminCollection(ctx, dbadmin.CollectionOrder)
	assert.NoError(t, err)
	servicemongo.Disconnect(ctx)
}
