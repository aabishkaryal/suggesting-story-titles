package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"

	"github.com/aabishkaryal/suggesting-story-titles/labels"
	"github.com/aabishkaryal/suggesting-story-titles/utils"
	"github.com/joho/godotenv"
)

const NUM_LABELS = 5

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

	labelCount := make(map[string]int)
	for _, rec := range records {
		recLabels := labels.LabelRecord(rec)
		for _, l := range recLabels {
			labelCount[l]++
		}
	}

	if len(labelCount) <= NUM_LABELS {
		for k := range labelCount {
			fmt.Println(k)
		}
		return
	}

	counts := make(sort.IntSlice, 0, len(labelCount))
	for _, v := range labelCount {
		counts = append(counts, v)
	}
	counts.Sort()
	for i, j := 0, len(counts)-1; i < j; i, j = i+1, j-1 {
		counts[i], counts[j] = counts[j], counts[i]
	}
	threshold := counts[NUM_LABELS]
	for k, v := range labelCount {
		if v >= threshold {
			fmt.Println(k)
		}
	}
}
