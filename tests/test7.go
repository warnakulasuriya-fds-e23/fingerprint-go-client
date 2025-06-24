package tests

import (
	"fmt"
	"log"
	"time"

	httprequests "github.com/warnakulasuriya-fds-e23/fingerprint-go-client/httpRequests"
)

func Test7() {
	fmt.Println("Identify endpoint test")
	client := httprequests.NewHttpClientImpl()
	var probeImagePath string = "/home/dheera/FingerPrintDatabases/veryLargePNGDataset/DB3_B107_1.png"

	t := time.Now()
	isMatched, discoveredId, err := client.IdentifyTemplateFilesMethod(probeImagePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	duration := time.Since(t)
	fmt.Printf("Is Matched: %t \n Discovered Id: %s\n", isMatched, discoveredId)
	fmt.Println("Elapsed Time: ", duration)
}
