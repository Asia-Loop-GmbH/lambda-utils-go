package service_common_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/logger"
	servicecommon "github.com/asia-loop-gmbh/lambda-utils-go/v3/service-common"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/test"
)

func TestGetGlobalConfig(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))
	log := logger.NewEmptyLogger()
	ctx := context.Background()

	cfg, err := servicecommon.GetGlobalConfig(log, ctx, "dev")
	Expect(err).To(BeNil())
	log.Infof("%v", cfg)
}
