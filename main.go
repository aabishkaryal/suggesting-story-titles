package main

import (
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
)

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

	for _, rec := range records {
		fmt.Println(labelRecord(rec))
	}

}

func labelRecord(record []string) string {
	lat, _ := strconv.ParseFloat(record[1], 32)
	lng, _ := strconv.ParseFloat(record[2], 32)
	date := parseDate(record[0])
	address := utils.ReverseGeocode(lat, lng)
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
