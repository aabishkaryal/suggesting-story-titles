package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	handleError(err, "Error loading environment variables")
}

func handleError(err error, message string) {
	if err != nil {
		log.Fatalln(message)
	}
}
