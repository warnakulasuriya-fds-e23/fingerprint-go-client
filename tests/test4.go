package tests

import (
	"fmt"
	"log"

	"github.com/warnakulasuriya-fds-e23/fingerprint-go-sdk/core"
)

func Test4() {
	var probeImagePath string = "/home/dheera/FingerPrintDatabases/veryLargePNGDataset/DB1_B101_1.png"
	var matchingImagePath string = probeImagePath
	var nonMatchingImagePath string = "/home/dheera/FingerPrintDatabases/veryLargePNGDataset/DB1_B102_1.png"

	var fingerprintImagesDirectoryPath string = "/home/dheera/FingerPrintDatabases/veryLargePNGDataset/"
	var cborDirectoryPath = "/home/dheera/FingerPrintDatabases/cborDirectory2ForGo/"

	sdk, err := core.NewsdkCore(fingerprintImagesDirectoryPath, cborDirectoryPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	probe := sdk.Extract(probeImagePath)
	matchingCandidate := sdk.Extract(matchingImagePath)
	nonMatchingCandidate := sdk.Extract(nonMatchingImagePath)

	fmt.Println("probe and matching Candidate: ", sdk.Match(probe, matchingCandidate))
	fmt.Println("probe and non matching Candidate: ", sdk.Match(probe, nonMatchingCandidate))

}
