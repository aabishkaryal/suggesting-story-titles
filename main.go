package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/aabishkaryal/suggesting-story-titles/labels"
	"github.com/aabishkaryal/suggesting-story-titles/utils"
	"github.com/joho/godotenv"
)

func main() {
	var fileName string
	if len(os.Args) < 2 {
		// log.Fatalln("Please provide a csv file to read data from.")
		fileName = "./data/0.csv"
	} else {
		fileName = os.Args[1]
	}

	err := godotenv.Load()
	utils.HandleError(err, "Error loading environment variables")

	utils.InitializeMapClient(os.Getenv("API_KEY"))

	file, err := os.Open(fileName)
	utils.HandleError(err, "Error opening file")

	records, err := csv.NewReader(file).ReadAll()
	utils.HandleError(err, "Error reading file")

	for _, rec := range records {
		fmt.Println(labels.LabelRecord(rec))
	}

}
