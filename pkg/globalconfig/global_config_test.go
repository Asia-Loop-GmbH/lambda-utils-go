package globalconfig_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/internal/pkg/test"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/globalconfig"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/logger"
)

func TestGetGlobalConfig(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))
	log := logger.NewEmptyLogger()
	ctx := context.Background()

	cfg, err := globalconfig.GetGlobalConfig(log, ctx, "dev")
	Expect(err).To(BeNil())
	log.Infof("%v", cfg)
}
