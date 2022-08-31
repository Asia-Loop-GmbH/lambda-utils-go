package servicecognito_test

import (
	"context"
	"log"
	"testing"

	commoncontext "github.com/nam-truong-le/lambda-utils-go/pkg/context"
	"github.com/stretchr/testify/assert"

	"github.com/asia-loop-gmbh/lambda-utils-go/v4/pkg/servicecognito"
)

func TestGetUser(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	ctx := context.WithValue(context.TODO(), commoncontext.FieldStage, "dev")
	user, err := servicecognito.GetUser(ctx, &servicecognito.GetUserData{
		Username: "lenamtruong@gmail.com",
	})
	assert.NoError(t, err)
	assert.Equal(t, "lenamtruong@gmail.com", user.Username)
	log.Printf("%v", user)
}
