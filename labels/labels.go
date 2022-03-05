package labels

import (
	"fmt"
	"strconv"
	"time"

	"github.com/aabishkaryal/suggesting-story-titles/models"
	"github.com/aabishkaryal/suggesting-story-titles/utils"
)

func LabelRecord(record []string) []string {
	lat, _ := strconv.ParseFloat(record[1], 32)
	lng, _ := strconv.ParseFloat(record[2], 32)
	date := utils.ParseDate(record[0])
	address := utils.ReverseGeocode(lat, lng)
	return generateLabels(date, address)
}

func generateLabels(date time.Time, address models.Address) []string {
	labels := make([]string, 0, 5)
	if date.Weekday() == time.Saturday || date.Weekday() == time.Sunday {
		if address.Locality != "" {
			labels = append(labels, fmt.Sprintf("A weekend in %s", address.Locality))
		}
		if address.City != "" {
			labels = append(labels, fmt.Sprintf("A weekend in %s", address.City))
		}
	}
	if address.City != "" {
		labels = append(labels, fmt.Sprintf("A trip to %s", address.City))
		labels = append(labels, fmt.Sprintf("%s in %s.", address.City, date.Month().String()))
	}
	if address.Country != "" {
		labels = append(labels, fmt.Sprintf("A trip to %s", address.Country))
	}
	return labels
}
