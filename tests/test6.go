package tests

import (
	"fmt"
	"log"
	"time"

	httprequests "github.com/warnakulasuriya-fds-e23/fingerprint-go-client/httpRequests"
)

func Test6(client *httprequests.Httpclientimpl) {
	fmt.Println("Match endpoint test")
	var probeImagePath string = "/home/dheera/FingerPrintDatabases/veryLargePNGDataset/DB1_B101_1.png"
	var matchingImagePath string = probeImagePath
	var nonMatchingImagePath string = "/home/dheera/FingerPrintDatabases/veryLargePNGDataset/DB1_B102_1.png"

	t := time.Now()
	isMatch1, err := client.MatchTemplatesFilesMethod(probeImagePath, matchingImagePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	firstduration := time.Since(t)

	t = time.Now()
	isMatch2, err := client.MatchTemplatesFilesMethod(probeImagePath, nonMatchingImagePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	secondduration := time.Since(t)

	fmt.Println("probe and matching Candidate: ", isMatch1)
	fmt.Println("probe and non matching Candidate: ", isMatch2)

	fmt.Println("Time elapsed for first match: ", firstduration)
	fmt.Println("Time elapsed for second match: ", secondduration)

}
