package servicecognito_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/internal/pkg/test"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/logger"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicecognito"
)

func TestGetUser(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))
	log := logger.NewEmptyLogger()

	user, err := servicecognito.GetUser(log, context.TODO(), &servicecognito.GetUserData{
		Username: "lenamtruong@gmail.com",
	})
	Expect(err).To(BeNil())
	Expect(user.Username).To(Equal("lenamtruong@gmail.com"))
	log.Infof("%v", user)
}
