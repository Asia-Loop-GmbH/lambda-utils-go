package globalconfig_test

import (
	"context"
	"testing"

	"github.com/nam-truong-le/lambda-utils-go/pkg/logger"
	"github.com/stretchr/testify/assert"

	"github.com/asia-loop-gmbh/lambda-utils-go/v4/pkg/globalconfig"
)

func TestGetGlobalConfig(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	ctx := context.Background()
	log := logger.FromContext(ctx)

	cfg, err := globalconfig.GetGlobalConfig(ctx)
	assert.NoError(t, err)
	log.Infof("%v", cfg)
}
