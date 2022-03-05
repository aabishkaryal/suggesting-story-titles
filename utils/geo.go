package utils

import (
	"context"

	"googlemaps.github.io/maps"
)

var mapsClient *maps.Client

func Initialize(apiKey string) {
	var err error
	mapsClient, err = maps.NewClient(maps.WithAPIKey(apiKey))
	HandleError(err, "Error initializing the google maps API client")
}

func ReverseGeocode(lat float64, lng float64) string {
	request := &maps.GeocodingRequest{LatLng: &maps.LatLng{Lat: lat, Lng: lng},
		ResultType: []string{"colloquial_area", "sublocality_level_1", "locality", "country"}}

	responses, err := mapsClient.ReverseGeocode(context.Background(), request)
	HandleError(err, "Error reverse geocoding lat and long")
	if len(responses) == 0 {
		return ""
	}
	return responses[0].FormattedAddress
}
