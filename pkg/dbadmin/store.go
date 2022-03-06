package dbadmin

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Store struct {
	ID            primitive.ObjectID `bson:"_id" json:"id"`
	Email         string             `bson:"email" json:"email"`
	Telephone     string             `bson:"telephone" json:"telephone"`
	Name          string             `bson:"name" json:"name"`
	Address       string             `bson:"address" json:"address"`
	Company       string             `bson:"company" json:"company"`
	Owner         string             `bson:"owner" json:"owner"`
	Register      string             `bson:"register" json:"register"`
	Tax           string             `bson:"tax" json:"tax"`
	Configuration StoreConfiguration `bson:"configuration" json:"configuration"`
	Drivers       []string           `bson:"drivers" json:"drivers"`
	Devices       []Device           `bson:"devices" json:"devices"`
	CreatedAt     time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt     time.Time          `bson:"updatedAt" json:"updatedAt"`
}

type StoreConfiguration struct {
	EnablePrinterAddress      bool   `bson:"enablePrinterAddress" json:"enablePrinterAddress"`
	EnablePrinterInternal     bool   `bson:"enablePrinterInternal" json:"enablePrinterInternal"`
	EnablePrinterPositions    bool   `bson:"enablePrinterPositions" json:"enablePrinterPositions"`
	POSID                     string `bson:"posId" json:"posId"`
	EnableAutomaticPosPayment bool   `bson:"enableAutomaticPosPayment" json:"enableAutomaticPosPayment"`
	WPStoreKey                string `bson:"wpStoreKey" json:"wpStoreKey"`
}

type Device struct {
	Name     string `bson:"name" json:"name"`
	DeviceID string `bson:"deviceId" json:"deviceId"`
}
