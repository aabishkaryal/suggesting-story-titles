package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

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
	date := utils.ParseDate(record[0])
	address := utils.ReverseGeocode(lat, lng)
	return date.Month().String() + "\t" + address
}
