package tests

import (
	"fmt"
	"log"
	"time"

	"github.com/warnakulasuriya-fds-e23/fingerprint-go-client/internal/httprequests"
)

func Test11(client *httprequests.Httpclientimpl) {
	fmt.Println("Authorize endpoint test")
	var probeImagePath string = "/home/dheera/FingerPrintDatabases/veryLargePNGDataset/FCV2000_DB1_B101_1.png"

	t := time.Now()
	status, err := client.AuthorizeTemplateFilesMethod(probeImagePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	duration := time.Since(t)
	fmt.Printf("Status: %s \n", status)
	fmt.Println("Elapsed Time: ", duration)
}
