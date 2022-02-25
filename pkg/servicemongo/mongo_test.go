package servicemongo_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/asia-loop-gmbh/lambda-types-go/v2/pkg/admin"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/internal/pkg/test"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/logger"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicemongo"
)

func TestAdminCollection(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))
	_, err := servicemongo.AdminCollection(logger.NewEmptyLogger(), context.Background(), "dev", admin.CollectionOrder)
	Expect(err).To(BeNil())
	_, err = servicemongo.AdminCollection(logger.NewEmptyLogger(), context.Background(), "dev", admin.CollectionOrder)
	Expect(err).To(BeNil())
	servicemongo.Disconnect(logger.NewEmptyLogger(), context.Background())
}
