package tests

import (
	"fmt"
	"log"
	"time"

	httprequests "github.com/warnakulasuriya-fds-e23/fingerprint-go-client/httpRequests"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-sdk/core"
)

func Test7() {
	fmt.Println("Identify endpoint test")
	client := httprequests.NewHttpClientImpl()
	var probeImagePath string = "/home/dheera/FingerPrintDatabases/veryLargePNGDataset/DB3_B107_1.png"

	var fingerprintImagesDirectoryPath string = "/home/dheera/FingerPrintDatabases/veryLargePNGDataset/"
	var cborDirectoryPath = "/home/dheera/FingerPrintDatabases/cborDirectory2ForGo/"

	sdk, err := core.NewSDKCore(fingerprintImagesDirectoryPath, cborDirectoryPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	probe, err := sdk.Extract(probeImagePath)
	if err != nil {
		log.Fatal(err.Error())
	}

	t := time.Now()
	isMatched, discoveredId := client.IdentifyTemplate(probe)
	duration := time.Since(t)
	fmt.Printf("Is Matched: %t \n Discovered Id: %s\n", isMatched, discoveredId)
	fmt.Println("Elapsed Time: ", duration)
}
