package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-client/internal/incoming"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-client/internal/outgoing"
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
	client := outgoing.NewOutgoingClient()

	// tests.Test11(client)
	// tests.Test9()
	// uploader.Uploader(client)
	IncomingController := incoming.NewIncomingController(*client)
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/api/file-submit", IncomingController.FileSubmitHandler)
	router.Run(":3000")
}
