package tests

import (
	"fmt"
	"log"
	"time"

	"github.com/warnakulasuriya-fds-e23/biometric-interface-service/internal/outgoing"
)

func Test11(client *outgoing.OutgoingClient) {
	fmt.Println("Authorize endpoint test")
	var probeImagePath string = "/home/dheera/FingerPrintDatabases/veryLargePNGDataset/FCV2004_DB1_B103_1.png"

	t := time.Now()
	status, err := client.AuthorizeTemplateFilesMethod(probeImagePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	duration := time.Since(t)
	fmt.Printf("Status: %s \n", status)
	fmt.Println("Elapsed Time: ", duration)
}
