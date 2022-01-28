package service_order

import (
	"context"
	"fmt"
	"github.com/asia-loop-gmbh/lambda-types-go/admin"
	"github.com/asia-loop-gmbh/lambda-utils-go/address"
	"github.com/asia-loop-gmbh/lambda-utils-go/mongo"
	"github.com/asia-loop-gmbh/lambda-utils-go/normalizer"
	"github.com/asia-loop-gmbh/lambda-utils-go/text"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func CreateOrder(stage string, orderOptions *admin.CreateOrderOrderOptions,
	addressOption *admin.CreateOrderAddressOptions) (*admin.Order, error) {
	log.Printf("create order: %s", orderOptions.OrderID)

	firstName := normalizer.Name(addressOption.FirstName)
	lastName := normalizer.Name(addressOption.LastName)
	telephone := normalizer.PhoneNumber(addressOption.Telephone)
	email := normalizer.Email(addressOption.Email)

	var addressLine1 string
	addressLine2 := addressOption.AddressLine2
	var postcode string
	var city string
	var formattedAddress string
	var validAddress bool

	inputAddress := fmt.Sprintf(
		"%s, %s %s",
		addressOption.AddressLine1,
		addressOption.Postcode,
		addressOption.City,
	)

	resolveAddressResult, err := address.ResolveAddress(inputAddress)
	if err != nil {
		log.Printf("could not resolve address: %s", inputAddress)
		addressLine1 = addressOption.AddressLine1
		postcode = addressOption.Postcode
		city = addressOption.City
		formattedAddress = inputAddress
		validAddress = false
	} else {
		addressLine1 = fmt.Sprintf("%s %s", resolveAddressResult.Street, resolveAddressResult.StreetNumber)
		postcode = resolveAddressResult.Postcode
		city = resolveAddressResult.City
		formattedAddress = resolveAddressResult.FormattedAddress
		validAddress = true
	}

	client, database, err := mongo.NewMongoAdminClient(context.Background(), stage)
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())

	collectionCustomer := client.Database(*database).Collection(admin.CollectionCustomer)
	findCustomer := collectionCustomer.FindOne(context.Background(), bson.M{
		"addressLine1": addressLine1,
		"addressLine2": addressLine2,
		"postcode":     postcode,
		"city":         city,
		"firstName":    firstName,
		"lastName":     lastName,
	})
	customer := new(admin.Customer)
	err = findCustomer.Decode(customer)
	if err != nil {
		log.Printf("could not find customer: %s %s (%s)", firstName, lastName, formattedAddress)
		customerRef := fmt.Sprintf(
			"%s%s",
			text.RandomString(2, false, true, false),
			text.RandomString(6, false, false, true),
		)
		newCustomer := admin.Customer{
			ID:           primitive.NewObjectID(),
			FirstName:    firstName,
			LastName:     lastName,
			CustomerRef:  customerRef,
			Boxes:        []int{},
			AddressLine1: addressLine1,
			AddressLine2: addressLine2,
			Postcode:     postcode,
			City:         city,
			Telephone:    telephone,
			Email:        email,
		}
		if _, err := collectionCustomer.InsertOne(context.Background(), newCustomer); err != nil {
			return nil, err
		}
		log.Printf("new customer created: %s %s (%s)", firstName, lastName, formattedAddress)
		customer = &newCustomer
	}

	collectionOrder := client.Database(*database).Collection(admin.CollectionOrder)
	secret := text.RandomString(32, true, false, true)
	newOrder := admin.Order{
		ID:                  primitive.NewObjectID(),
		Status:              orderOptions.Status,
		OrderID:             orderOptions.OrderID,
		OrderNumber:         orderOptions.OrderNumber,
		InvoiceNumber:       orderOptions.InvoiceNumber,
		ValidAddress:        validAddress,
		Address:             formattedAddress,
		AddressLine2:        addressLine2,
		Email:               email,
		Telephone:           telephone,
		CustomerNote:        orderOptions.CustomerNote,
		DeliveryDate:        orderOptions.DeliveryDate,
		DeliveryTime:        orderOptions.DeliveryTime,
		ShippingMethod:      orderOptions.ShippingMethod,
		Customer:            customer.ID,
		Items:               orderOptions.Items,
		Tip:                 orderOptions.Tip,
		Total:               orderOptions.Total,
		Tax:                 orderOptions.Tax,
		Net:                 orderOptions.Net,
		CouponCode:          orderOptions.CouponCode,
		AppliedCouponTax:    orderOptions.AppliedCouponTax,
		AppliedCouponTotal:  orderOptions.AppliedCouponTotal,
		AppliedCouponNet:    orderOptions.AppliedCouponNet,
		Store:               orderOptions.Store,
		CompanyKey:          orderOptions.CompanyKey,
		Secret:              secret,
		Printed:             false,
		LocalPickupNotified: false,
		PaymentEvents:       make([]interface{}, 0),
	}

	if _, err := collectionOrder.InsertOne(context.Background(), newOrder); err != nil {
		return nil, err
	}

	log.Printf("order created: %s", orderOptions.OrderID)
	if customer.Telephone == "" {
		_, err := collectionOrder.UpdateByID(context.Background(), customer.ID, bson.M{
			"telephone": telephone,
		})
		if err != nil {
			log.Printf("could not update customer telphone to: %s", telephone)
		} else {
			log.Printf("customer telphone updated: %s", telephone)
		}
	}
	_, err = collectionOrder.UpdateByID(context.Background(), customer.ID, bson.M{
		"email": email,
	})
	if err != nil {
		log.Printf("could not update customer email to: %s", email)
	} else {
		log.Printf("customer email updated: %s", email)
	}

	// TODO: post processing

	return &newOrder, nil
}
