package servicecognito_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/internal/pkg/test"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/logger"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicecognito"
)

func TestCreateUser(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))
	log := logger.NewEmptyLogger()
	ctx := context.TODO()

	err := servicecognito.DeleteUser(log, ctx, &servicecognito.DeleteUserData{
		Username: "lenamtruong+unitest@gmail.com",
	})
	if err != nil {
		log.Infof("it's ok, just a cleanup")
	}

	err = servicecognito.CreateUser(log, ctx, &servicecognito.CreateUserData{
		Username:          "lenamtruong+unitest@gmail.com",
		TemporaryPassword: "Toor***???1234",
		CompanyKey:        "SIEMENS",
		StoreKey:          "ERLANGEN",
		FirstName:         "Nam Truong",
		LastName:          "Le",
		Role:              "COMPANY",
	})
	Expect(err).To(BeNil())

	user, err := servicecognito.GetUser(log, ctx, &servicecognito.GetUserData{
		Username: "lenamtruong+unitest@gmail.com",
	})
	Expect(err).To(BeNil())
	Expect(user.Username).To(Equal("lenamtruong+unitest@gmail.com"))
	Expect(user.FirstName).To(Equal("Nam Truong"))
	Expect(user.LastName).To(Equal("Le"))
	Expect(user.Company).To(Equal("SIEMENS"))

	err = servicecognito.DeleteUser(log, ctx, &servicecognito.DeleteUserData{
		Username: "lenamtruong+unitest@gmail.com",
	})
	Expect(err).To(BeNil())

	_, err = servicecognito.GetUser(log, ctx, &servicecognito.GetUserData{
		Username: "lenamtruong+unitest@gmail.com",
	})
	Expect(err).ToNot(BeNil())
}
