package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
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
	// client := httprequests.NewHttpClientImpl()
	// config := configtomlreader.ConfigTomlReader()
	// client.SetOrAddHeaderValueAccordingToKey("Test-Key", config.TestKey)

	tests.Test9()
	// uploader.Uploader(client)
}
