package servicepusher

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
)

const (
	eventOrderCreatedName           = "order-created"
	eventOrderPOSPaymentStartedName = "order-payment-started"
	eventOrderPOSPaymentPaidName    = "order-paid"
	eventOrderDeliveredName         = "order-delivered"
	eventGroupFinalizedName         = "group-finalized"
	eventGroupDeliveredName         = "group-delivered"
)

type EventOrderCreatedData struct {
	StoreID string `json:"storeId"`
}

func PublishOrderCreated(log *logrus.Entry, ctx context.Context, stage string, data *EventOrderCreatedData) error {
	client, err := getClient(log, ctx, stage)
	if err != nil {
		return err
	}
	if err := client.Trigger(storeChannel(data.StoreID), eventOrderCreatedName, data); err != nil {
		log.Errorf("failed to publish event [%s]: %v", eventOrderCreatedName, data)
		return err
	}
	log.Infof("event published [%s]: %v", eventOrderCreatedName, data)
	return nil
}

type EventOrderDeliveredData struct {
	StoreID string `json:"storeId"`
}

func PublishOrderDelivered(log *logrus.Entry, ctx context.Context, stage string, data *EventOrderDeliveredData) error {
	client, err := getClient(log, ctx, stage)
	if err != nil {
		return err
	}
	if err := client.Trigger(storeChannel(data.StoreID), eventOrderDeliveredName, data); err != nil {
		log.Errorf("failed to publish event [%s]: %v", eventOrderDeliveredName, data)
		return err
	}
	log.Infof("event published [%s]: %v", eventOrderDeliveredName, data)
	return nil
}

type EventOrderPOSPaymentStartedData struct {
	StoreID string `json:"storeId"`
	OrderID string `json:"orderId"`
}

func PublishOrderPOSPaymentStarted(log *logrus.Entry, ctx context.Context, stage string, data *EventOrderPOSPaymentStartedData) error {
	client, err := getClient(log, ctx, stage)
	if err != nil {
		return err
	}
	if err := client.Trigger(storeChannel(data.StoreID), eventOrderPOSPaymentStartedName, data); err != nil {
		log.Errorf("failed to publish event [%s]: %v", eventOrderPOSPaymentStartedName, data)
		return err
	}
	log.Infof("event published [%s]: %v", eventOrderPOSPaymentStartedName, data)
	return nil
}

type EventOrderPOSPaymentPaidData struct {
	StoreID string `json:"storeId"`
	OrderID string `json:"orderId"`
}

func PublishOrderPOSPaymentPaid(log *logrus.Entry, ctx context.Context, stage string, data *EventOrderPOSPaymentPaidData) error {
	client, err := getClient(log, ctx, stage)
	if err != nil {
		return err
	}
	if err := client.Trigger(storeChannel(data.StoreID), eventOrderPOSPaymentPaidName, data); err != nil {
		log.Errorf("failed to publish event [%s]: %v", eventOrderPOSPaymentPaidName, data)
		return err
	}
	log.Infof("event published [%s]: %v", eventOrderPOSPaymentPaidName, data)
	return nil
}

type EventGroupFinalizedData struct {
	StoreID string `json:"storeId"`
}

func PublishGroupFinalized(log *logrus.Entry, ctx context.Context, stage string, data *EventGroupFinalizedData) error {
	client, err := getClient(log, ctx, stage)
	if err != nil {
		return err
	}
	if err := client.Trigger(storeChannel(data.StoreID), eventGroupFinalizedName, data); err != nil {
		log.Errorf("failed to publish event [%s]: %v", eventGroupFinalizedName, data)
		return err
	}
	log.Infof("event published [%s]: %v", eventGroupFinalizedName, data)
	return nil
}

type EventGroupDeliveredData struct {
	StoreID string `json:"storeId"`
}

func PublishGroupDelivered(log *logrus.Entry, ctx context.Context, stage string, data *EventGroupDeliveredData) error {
	client, err := getClient(log, ctx, stage)
	if err != nil {
		return err
	}
	if err := client.Trigger(storeChannel(data.StoreID), eventGroupDeliveredName, data); err != nil {
		log.Errorf("failed to publish event [%s]: %v", eventGroupDeliveredName, data)
		return err
	}
	log.Infof("event published [%s]: %v", eventGroupDeliveredName, data)
	return nil
}

func storeChannel(storeID string) string {
	return fmt.Sprintf("restaurant-%s", storeID)
}
