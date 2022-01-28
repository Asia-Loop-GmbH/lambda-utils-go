package address

import (
	"context"
	"fmt"
	utils "github.com/asia-loop-gmbh/lambda-utils-go/aws"
	"googlemaps.github.io/maps"
	"log"
)

type ResolveAddressResult struct {
	StreetNumber     string
	Street           string
	City             string
	Postcode         string
	State            string
	FormattedAddress string
}

func ResolveAddress(address string) (*ResolveAddressResult, error) {
	apiKey, err := utils.GetSSMParameter("all", "/google/maps/key", true)
	if err != nil {
		return nil, err
	}
	client, err := maps.NewClient(maps.WithAPIKey(*apiKey))
	if err != nil {
		return nil, err
	}
	result, err := client.Geocode(context.Background(), &maps.GeocodingRequest{
		Address:  address,
		Region:   "DE",
		Language: "DE",
	})
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("could not resolve address: %s", address)
	}

	if len(result) > 1 {
		log.Printf("multiple addresses found, the first item will be taken")
	}

	return &ResolveAddressResult{
		StreetNumber:     getAddressComponent(result[0], "street_number"),
		Street:           getAddressComponent(result[0], "route"),
		City:             getAddressComponent(result[0], "locality"),
		Postcode:         getAddressComponent(result[0], "postal_code"),
		State:            getAddressComponent(result[0], "administrative_area_level_1"),
		FormattedAddress: result[0].FormattedAddress,
	}, nil
}

func getAddressComponent(result maps.GeocodingResult, componentName string) string {
	for _, component := range result.AddressComponents {
		for _, componentType := range component.Types {
			if componentType == componentName {
				return component.LongName
			}
		}
	}
	return ""
}
