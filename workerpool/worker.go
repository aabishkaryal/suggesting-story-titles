package workerpool

import "github.com/aabishkaryal/suggesting-story-titles/labels"

func Worker(jobs <-chan []string, results chan<- []string) {
	for job := range jobs {
		results <- labels.LabelRecord(job)
	}
}
