package tests

import (
	"fmt"
	"log"
	"time"

	httprequests "github.com/warnakulasuriya-fds-e23/fingerprint-go-client/internal/httprequests"
)

func Test7(client *httprequests.Httpclientimpl) {
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
