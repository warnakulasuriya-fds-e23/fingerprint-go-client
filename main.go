package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	httprequests "github.com/warnakulasuriya-fds-e23/fingerprint-go-client/httpRequests"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-client/tests"
)

func main() {
	_, err := os.Stat(".env")
	if err == nil {
		log.Println("discovered .env file")
		err := godotenv.Load()
		if err != nil {
			log.Println("however failed to load .env file")
		} else {
			log.Println(".env successfully loaded")
		}
	}
	client := httprequests.NewHttpClientImpl()

	tests.Test7(client)
	// uploader.Uploader(client)
}
