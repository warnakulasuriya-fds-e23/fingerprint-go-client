package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-client/internal/httprequests"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-client/internal/tests"
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
	//TODO: try to utilize github.com/tidwall/gjson"
	client := httprequests.NewHttpClientImpl()

	tests.Test7(client)
	// uploader.Uploader(client)
}
