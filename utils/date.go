package utils

import (
	"errors"
	"regexp"
	"time"
)

func ParseDate(date string) time.Time {
	if match, _ := regexp.MatchString(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z$`, date); match {
		t, err := time.Parse(time.RFC3339, date)
		HandleError(err, "Error parsing date")
		return t
	} else if match, _ = regexp.MatchString(`\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}`, date); match {
		t, err := time.Parse("2006-01-02 15:04:05", date)
		HandleError(err, "Error parsing date ")
		return t
	}
	HandleError(errors.New("unsupported date format"), "Error parsing date "+date)
	return time.Time{}
}
