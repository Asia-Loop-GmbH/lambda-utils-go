package servicesns_test

import (
	"context"
	"testing"

	commoncontext "github.com/nam-truong-le/lambda-utils-go/pkg/context"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/asia-loop-gmbh/lambda-utils-go/v4/pkg/servicesns"
)

func TestPublishOrderFinalized(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	ctx := context.WithValue(context.TODO(), commoncontext.FieldStage, "dev")
	id, _ := primitive.ObjectIDFromHex("622a5275b73e4d6262fd8acf")
	err := servicesns.PublishOrderFinalized(ctx, &servicesns.EventOrderFinalizedData{
		ID: id,
	})
	assert.NoError(t, err)
}
