package utils

import (
	"context"
	"log"

	"github.com/aabishkaryal/suggesting-story-titles/models"
	"googlemaps.github.io/maps"
)

var mapClient *maps.Client

func InitializeMapClient(apiKey string) {
	var err error
	mapClient, err = maps.NewClient(maps.WithAPIKey(apiKey))
	HandleError(err, "Error initializing the google maps API client")
}

func ReverseGeocode(lat float64, lng float64) models.Address {
	request := &maps.GeocodingRequest{LatLng: &maps.LatLng{Lat: lat, Lng: lng},
		ResultType: []string{"colloquial_area", "sublocality_level_1", "locality", "country"}}

	responses, err := mapClient.ReverseGeocode(context.Background(), request)
	HandleError(err, "Error reverse geocoding lat and long")
	if len(responses) == 0 {
		log.Fatalln("Unable to reverse geocode")
	}

	address := models.Address{}
	for _, addressComponent := range responses[0].AddressComponents {
		if Contains(addressComponent.Types, "country") {
			address.Country = addressComponent.LongName
		} else if Contains(addressComponent.Types, "locality") {
			address.City = addressComponent.LongName
		} else if Contains(addressComponent.Types, "sublocality") {
			address.Locality = addressComponent.LongName
		}
	}
	return address
}
