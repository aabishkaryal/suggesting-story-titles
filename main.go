package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"

	"github.com/aabishkaryal/suggesting-story-titles/utils"
	"github.com/aabishkaryal/suggesting-story-titles/workerpool"
	"github.com/joho/godotenv"
)

// Show the top 3 most popular Labels
// Could be more if multiple labels share the same position
const NUM_LABELS = 3

// Number of workers/goroutines
const NUM_WORKER = 8

func main() {
	var fileName string
	if len(os.Args) < 2 {
		// log.Fatalln("Please provide a csv file to read data from.")
		fileName = "./data/0.csv"
	} else {
		fileName = os.Args[1]
	}

	// Load env variables with godotenv
	err := godotenv.Load()
	utils.HandleError(err, "Error loading environment variables")

	// Initialize the map client
	utils.InitializeMapClient(os.Getenv("API_KEY"))

	// Open and read the file
	file, err := os.Open(fileName)
	utils.HandleError(err, "Error opening file")
	records, err := csv.NewReader(file).ReadAll()
	utils.HandleError(err, "Error reading file")

	// Create channels for concurrency
	jobs := make(chan []string, len(records))
	results := make(chan []string, len(records))

	// Spawn workers
	for i := 0; i < NUM_WORKER; i++ {
		go workerpool.Worker(jobs, results)
	}

	for _, rec := range records {
		jobs <- rec
	}
	close(jobs)

	// Keep track of how many times a label was recommended
	labelCount := make(map[string]int)
	for range records {
		recLabels := <-results
		for _, l := range recLabels {
			labelCount[l]++
		}
	}

	// Print all labels if number of all labels is lower than how many we want
	if len(labelCount) <= NUM_LABELS {
		for k := range labelCount {
			fmt.Println(k)
		}
		return
	}

	// sort the count to decide the threshold to ignore labels
	counts := make(sort.IntSlice, 0, len(labelCount))
	for _, v := range labelCount {
		counts = append(counts, v)
	}
	counts.Sort()
	for i, j := 0, len(counts)-1; i < j; i, j = i+1, j-1 {
		counts[i], counts[j] = counts[j], counts[i]
	}

	threshold := counts[NUM_LABELS]
	// Only show labels with recommendation above the threshold
	for k, v := range labelCount {
		if v >= threshold {
			fmt.Println(k)
		}
	}
}
