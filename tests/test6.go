package tests

import (
	"fmt"
	"log"
	"time"

	httprequests "github.com/warnakulasuriya-fds-e23/fingerprint-go-client/httpRequests"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-sdk/core"
)

func Test6() {
	client := httprequests.NewHttpClientImpl()
	var probeImagePath string = "/home/dheera/FingerPrintDatabases/veryLargePNGDataset/DB1_B101_1.png"
	var matchingImagePath string = probeImagePath
	var nonMatchingImagePath string = "/home/dheera/FingerPrintDatabases/veryLargePNGDataset/DB1_B102_1.png"

	var fingerprintImagesDirectoryPath string = "/home/dheera/FingerPrintDatabases/veryLargePNGDataset/"
	var cborDirectoryPath = "/home/dheera/FingerPrintDatabases/cborDirectory2ForGo/"

	sdk, err := core.NewSDKCore(fingerprintImagesDirectoryPath, cborDirectoryPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	probe := sdk.Extract(probeImagePath)
	matchingCandidate := sdk.Extract(matchingImagePath)
	nonMatchingCandidate := sdk.Extract(nonMatchingImagePath)

	t := time.Now()
	isMatch1 := client.MatchTemplates(probe, matchingCandidate)
	firstduration := time.Since(t)

	t = time.Now()
	isMatch2 := client.MatchTemplates(probe, nonMatchingCandidate)
	secondduration := time.Since(t)

	fmt.Println("probe and matching Candidate: ", isMatch1)
	fmt.Println("probe and non matching Candidate: ", isMatch2)

	fmt.Println("Time elapsed for first match: ", firstduration)
	fmt.Println("Time elapsed for second match: ", secondduration)

}
