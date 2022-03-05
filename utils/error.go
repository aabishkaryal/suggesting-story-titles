package utils

import "log"

func HandleError(err error, message string) {
	if err != nil {
		log.Fatalf("%s \t %v", message, err)
	}
}
