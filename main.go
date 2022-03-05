package main

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/aabishkaryal/suggesting-story-titles/utils"
	"github.com/joho/godotenv"
	"googlemaps.github.io/maps"
)

var mapsClient *maps.Client

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Please provide a csv file to read data from.")
	}
	file, err := os.Open(os.Args[1])
	utils.HandleError(err, "Error opening file")

	records, err := csv.NewReader(file).ReadAll()
	utils.HandleError(err, "Error reading file")

	err = godotenv.Load()
	utils.HandleError(err, "Error loading environment variables")

	mapsClient, err = maps.NewClient(maps.WithAPIKey(os.Getenv("API_KEY")))
	utils.HandleError(err, "Error initializing the google maps API client")

	for _, rec := range records {
		fmt.Println(labelRecord(rec))
	}

}

func labelRecord(record []string) string {
	lat, _ := strconv.ParseFloat(record[1], 32)
	lng, _ := strconv.ParseFloat(record[2], 32)
	date := parseDate(record[0])
	address := reverseGeocode(lat, lng)
	return date.Month().String() + "\t" + address
}

func parseDate(date string) time.Time {
	if match, _ := regexp.MatchString(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z$`, date); match {
		t, err := time.Parse(time.RFC3339, date)
		utils.HandleError(err, "Error parsing date")
		return t
	} else if match, _ = regexp.MatchString(`\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}`, date); match {
		t, err := time.Parse("2006-01-02 15:04:05", date)
		utils.HandleError(err, "Error parsing date ")
		return t
	}
	utils.HandleError(errors.New("unsupported date format"), "Error parsing date "+date)
	return time.Time{}
}

func reverseGeocode(lat float64, lng float64) string {
	request := &maps.GeocodingRequest{LatLng: &maps.LatLng{Lat: lat, Lng: lng},
		ResultType: []string{"colloquial_area", "sublocality_level_1", "locality", "country"}}

	responses, err := mapsClient.ReverseGeocode(context.Background(), request)
	utils.HandleError(err, "Error reverse geocoding lat and long")
	if len(responses) == 0 {
		return ""
	}
	return responses[0].FormattedAddress
}
