package servicessm_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/asia-loop-gmbh/lambda-utils-go/v4/pkg/servicessm"
)

func TestGetParameter_FromEnv(t *testing.T) {
	dummyValue := "foo"
	err := os.Setenv("AL_SSM_ALL_MONGO_HOST", dummyValue)
	assert.NoError(t, err)
	v, err := servicessm.GetGlobalParameter(context.TODO(), "/mongo/host", false)

	assert.Equal(t, dummyValue, *v)
}

func TestGetParameter_EnvNotExists(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	err := os.Unsetenv("AL_SSM_ALL_MONGO_HOST")
	assert.NoError(t, err)
	v, err := servicessm.GetGlobalParameter(context.TODO(), "/mongo/host", false)

	assert.Equal(t, "asia-loop-admin.yncuk.mongodb.net", *v)
}

func TestGetParameter_EnvEmptyString(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	err := os.Setenv("AL_SSM_ALL_MONGO_HOST", "")
	assert.NoError(t, err)
	v, err := servicessm.GetGlobalParameter(context.TODO(), "/mongo/host", false)

	assert.Equal(t, "asia-loop-admin.yncuk.mongodb.net", *v)
}
