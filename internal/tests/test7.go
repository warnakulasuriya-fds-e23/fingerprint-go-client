package tests

import (
	"fmt"
	"log"
	"time"

	"github.com/warnakulasuriya-fds-e23/biometric-interface-service/internal/outgoing"
)

func Test7(client *outgoing.OutgoingClient) {
	fmt.Println("Identify endpoint test")
	var probeImagePath string = "/home/dheera/FingerPrintDatabases/veryLargePNGDataset/FCV2000_DB1_B101_1.png"

	t := time.Now()
	isMatched, discoveredId, err := client.IdentifyTemplateFilesMethod(probeImagePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	duration := time.Since(t)
	fmt.Printf("Is Matched: %t \n Discovered Id: %s\n", isMatched, discoveredId)
	fmt.Println("Elapsed Time: ", duration)
}
