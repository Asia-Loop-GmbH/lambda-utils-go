package coupon_test

import (
	"context"
	"testing"

	commoncontext "github.com/nam-truong-le/lambda-utils-go/pkg/context"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/asia-loop-gmbh/lambda-utils-go/v4/pkg/servicewoo/coupon"
)

func TestGetCouponByCode_Success(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	ctx := context.WithValue(context.TODO(), commoncontext.FieldStage, "dev")
	c, err := coupon.GetCouponByCode(ctx, "TEST_COUPON")
	assert.NoError(t, err)
	assert.Equal(t, "test_coupon", c.Code)
}

func TestGetCouponByCode_Fail(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	ctx := context.WithValue(context.TODO(), commoncontext.FieldStage, "dev")
	_, err := coupon.GetCouponByCode(ctx, "TEST_COUPON_NOT_EXISTS")
	assert.NoError(t, err)
}

func TestIsValidAndHasEnough_Success(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	ctx := context.WithValue(context.TODO(), commoncontext.FieldStage, "dev")
	valid := coupon.IsValidAndHasEnough(ctx, "TEST_COUPON", "10.00")
	assert.True(t, valid)
}

func TestIsValidAndHasEnough_Fail(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	ctx := context.WithValue(context.TODO(), commoncontext.FieldStage, "dev")
	valid := coupon.IsValidAndHasEnough(ctx, "TEST_COUPON", "10000.00")
	assert.False(t, valid)
}

func TestUpdateCouponByCode_Success(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	ctx := context.WithValue(context.TODO(), commoncontext.FieldStage, "dev")
	amount := "0.01"

	c, err := coupon.GetCouponByCode(ctx, "TEST_COUPON")
	assert.NoError(t, err)

	current, err := decimal.NewFromString(c.Amount)
	assert.NoError(t, err)

	err = coupon.UpdateCouponByCode(ctx, "TEST_COUPON", amount)
	assert.NoError(t, err)

	c, err = coupon.GetCouponByCode(ctx, "TEST_COUPON")
	assert.NoError(t, err)

	updated, err := decimal.NewFromString(c.Amount)
	assert.NoError(t, err)

	assert.Equal(t, amount, current.Sub(updated).StringFixed(2))
}
