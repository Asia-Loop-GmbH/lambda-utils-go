package coupon_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/shopspring/decimal"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/internal/pkg/test"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/logger"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicewoo/coupon"
)

func TestGetCouponByCode_Success(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	c, err := coupon.GetCouponByCode(logger.NewEmptyLogger(), context.TODO(), "dev", "TEST_COUPON")
	Expect(err).To(BeNil())
	Expect(c.Code).To(Equal("test_coupon"))
}

func TestGetCouponByCode_Fail(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	_, err := coupon.GetCouponByCode(logger.NewEmptyLogger(), context.TODO(), "dev", "TEST_COUPON_NOT_EXISTS")
	Expect(err).To(Not(BeNil()))
}

func TestIsValidAndHasEnough_Success(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))
	valid := coupon.IsValidAndHasEnough(logger.NewEmptyLogger(), context.TODO(), "dev", "TEST_COUPON", "10.00")
	Expect(valid).To(BeTrue())
}

func TestIsValidAndHasEnough_Fail(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))
	valid := coupon.IsValidAndHasEnough(logger.NewEmptyLogger(), context.TODO(), "dev", "TEST_COUPON", "10000.00")
	Expect(valid).To(BeFalse())
}

func TestUpdateCouponByCode_Success(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))
	amount := "0.01"

	c, err := coupon.GetCouponByCode(logger.NewEmptyLogger(), context.TODO(), "dev", "TEST_COUPON")
	Expect(err).To(BeNil())

	current, err := decimal.NewFromString(c.Amount)
	Expect(err).To(BeNil())

	err = coupon.UpdateCouponByCode(logger.NewEmptyLogger(), context.TODO(), "dev", "TEST_COUPON", amount)
	Expect(err).To(BeNil())

	c, err = coupon.GetCouponByCode(logger.NewEmptyLogger(), context.TODO(), "dev", "TEST_COUPON")
	Expect(err).To(BeNil())

	updated, err := decimal.NewFromString(c.Amount)
	Expect(err).To(BeNil())

	Expect(current.Sub(updated).StringFixed(2)).To(Equal(amount))
}
