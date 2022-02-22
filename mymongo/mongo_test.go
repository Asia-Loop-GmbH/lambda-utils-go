package mymongo_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/asia-loop-gmbh/lambda-types-go/admin"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/logger"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/mymongo"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/test"
)

func TestAdminCollection(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))
	_, err := mymongo.AdminCollection(logger.NewEmptyLogger(), context.Background(), "dev", admin.CollectionOrder)
	Expect(err).To(BeNil())
	_, err = mymongo.AdminCollection(logger.NewEmptyLogger(), context.Background(), "dev", admin.CollectionOrder)
	Expect(err).To(BeNil())
	mymongo.Disconnect(logger.NewEmptyLogger(), context.Background())
}
