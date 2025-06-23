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

	sdk, err := core.NewSDKCore(fingerprintImagesDirectoryPath, cborDirectoryPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	probe, err := sdk.Extract(probeImagePath)
	if err != nil {
		log.Fatal(err.Error())
	}

	matchingCandidate, err := sdk.Extract(matchingImagePath)
	if err != nil {
		log.Fatal(err.Error())
	}

	nonMatchingCandidate, err := sdk.Extract(nonMatchingImagePath)
	if err != nil {
		log.Fatal(err.Error())
	}

	comparison1, err := sdk.Match(probe, matchingCandidate)
	if err != nil {
		log.Fatal(err.Error())
	}

	comparison2, err := sdk.Match(probe, nonMatchingCandidate)

	fmt.Println("probe and matching Candidate: ", comparison1)
	fmt.Println("probe and non matching Candidate: ", comparison2)

}
