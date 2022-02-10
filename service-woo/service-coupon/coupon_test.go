package service_coupon_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/shopspring/decimal"

	"github.com/asia-loop-gmbh/lambda-utils-go/logger"
	servicecoupon "github.com/asia-loop-gmbh/lambda-utils-go/service-woo/service-coupon"
	"github.com/asia-loop-gmbh/lambda-utils-go/test"
)

func TestGetCouponByCode_Success(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	coupon, err := servicecoupon.GetCouponByCode(logger.NewEmptyLogger(), "dev", "TEST_COUPON")
	Expect(err).To(BeNil())
	Expect(coupon.Code).To(Equal("test_coupon"))
}

func TestGetCouponByCode_Fail(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))

	_, err := servicecoupon.GetCouponByCode(logger.NewEmptyLogger(), "dev", "TEST_COUPON_NOT_EXISTS")
	Expect(err).To(Not(BeNil()))
}

func TestIsValidAndHasEnough_Success(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))
	valid := servicecoupon.IsValidAndHasEnough(logger.NewEmptyLogger(), "dev", "TEST_COUPON", "10.00")
	Expect(valid).To(BeTrue())
}

func TestIsValidAndHasEnough_Fail(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))
	valid := servicecoupon.IsValidAndHasEnough(logger.NewEmptyLogger(), "dev", "TEST_COUPON", "10000.00")
	Expect(valid).To(BeFalse())
}

func TestUpdateCouponByCode_Success(t *testing.T) {
	RegisterFailHandler(test.FailedHandler(t))
	amount := "0.01"

	coupon, err := servicecoupon.GetCouponByCode(logger.NewEmptyLogger(), "dev", "TEST_COUPON")
	Expect(err).To(BeNil())

	current, err := decimal.NewFromString(coupon.Amount)
	Expect(err).To(BeNil())

	err = servicecoupon.UpdateCouponByCode(logger.NewEmptyLogger(), "dev", "TEST_COUPON", amount)
	Expect(err).To(BeNil())

	coupon, err = servicecoupon.GetCouponByCode(logger.NewEmptyLogger(), "dev", "TEST_COUPON")
	Expect(err).To(BeNil())

	updated, err := decimal.NewFromString(coupon.Amount)
	Expect(err).To(BeNil())

	Expect(current.Sub(updated).StringFixed(2)).To(Equal(amount))
}
