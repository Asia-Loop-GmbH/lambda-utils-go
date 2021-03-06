package revenue

import (
	"context"
	"strconv"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/dbadmin"
	"github.com/asia-loop-gmbh/lambda-utils-go/v3/pkg/servicemongo"
)

func CheckRefund(log *logrus.Entry, ctx context.Context, stage, merchantRef string) error {
	synced, err := refundExists(log, ctx, stage, merchantRef)
	if err != nil {
		return err
	}

	if !synced {
		// not synchronized
		_, err := primitive.ObjectIDFromHex(merchantRef)
		isWoo := err != nil

		if isWoo {
			// woo refund
			wooID, err := strconv.Atoi(merchantRef)
			if err != nil {
				return err
			}
			return SyncWooRefunds(log, ctx, stage, wooID)
		}

		// pos/corporate refund
		return fakeNonWooRefund(log, ctx, stage, merchantRef)
	}

	// already synchronized
	return nil
}

func CheckOrder(log *logrus.Entry, ctx context.Context, stage, merchantRef string) error {
	refAsObjectID, err := primitive.ObjectIDFromHex(merchantRef)
	isWoo := err != nil

	if isWoo {
		synced, err := orderExists(log, ctx, stage, merchantRef)
		if err != nil {
			return err
		}
		if !synced {
			wooID, err := strconv.Atoi(merchantRef)
			if err != nil {
				return err
			}
			return SyncWooOrder(log, ctx, stage, wooID)
		}
		return nil
	}

	synced, err := orderExists(log, ctx, stage, merchantRef)
	if err != nil {
		return err
	}

	if !synced {
		colOrder, err := servicemongo.AdminCollection(log, ctx, stage, dbadmin.CollectionOrder)
		if err != nil {
			return err
		}
		findOrder := colOrder.FindOne(ctx, bson.M{"_id": refAsObjectID})
		o := new(dbadmin.Order)
		err = findOrder.Decode(o)
		if err != nil {
			return err
		}

		if o.CompanyKey != "" {
			return syncCorporateOrder(log, ctx, stage, o)
		} else {
			return syncPOSOrder(log, ctx, stage, o)
		}
	}
	return nil
}
