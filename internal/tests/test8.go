package tests

import (
	"fmt"
	"log"
	"time"

	"github.com/warnakulasuriya-fds-e23/fingerprint-go-client/internal/outgoing"
)

func Test8(client *outgoing.OutgoingClient) {
	fmt.Println("Enroll endpoint test")
	var NewEntryImagePath string = "/home/dheera/FingerPrintDatabases/veryLargePNGDataset/DB3_B107_1.png"
	t := time.Now()
	message, err := client.EnrollTemplateFilesMethod(NewEntryImagePath, "testuser1")
	duration := time.Since(t)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(message)
	fmt.Println("elapsed time : ", duration)
}
