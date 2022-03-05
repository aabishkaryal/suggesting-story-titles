package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"googlemaps.github.io/maps"
)

var mapsClient *maps.Client

func main() {
	err := godotenv.Load()
	handleError(err, "Error loading environment variables")

	handleError(err, "Error initializing google maps")
	mapsClient, err = maps.NewClient(maps.WithAPIKey(os.Getenv("API_KEY")))
	handleError(err, "Error initializing the google maps API client")

	fmt.Println(reverseGeocode(40.705882, -74.010093))
	fmt.Println(reverseGeocode(36.126489, -115.166550))
	fmt.Println(reverseGeocode(40.633113, 14.483203))
}

func reverseGeocode(lat float64, lng float64) string {
	request := &maps.GeocodingRequest{LatLng: &maps.LatLng{Lat: lat, Lng: lng},
		ResultType: []string{"colloquial_area", "sublocality_level_1", "locality", "country"}}

	responses, err := mapsClient.ReverseGeocode(context.Background(), request)
	handleError(err, "Error reverse geocoding lat and long")
	if len(responses) == 0 {
		return ""
	}
	return responses[0].FormattedAddress
}

func handleError(err error, message string) {
	if err != nil {
		log.Fatalln(message)
	}
}
