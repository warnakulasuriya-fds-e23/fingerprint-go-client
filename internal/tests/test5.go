package tests

import (
	"fmt"
	"log"

	"github.com/warnakulasuriya-fds-e23/fingerprint-go-sdk/core"
)

func Test5() {
	var fingerprintImagesDirectoryPath string = "/home/dheera/FingerPrintDatabases/smallPNGDataset/"
	var cborDirectoryPath = "/home/dheera/FingerPrintDatabases/cborDirectory2ForGo/"
	sdk, err := core.NewSDKCore(fingerprintImagesDirectoryPath, cborDirectoryPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	sdk.LoadImages()

	var probeImagePath string = "/home/dheera/FingerPrintDatabases/veryLargePNGDataset/DB3_B107_1.png"
	probe, err := sdk.Extract(probeImagePath)
	if err != nil {
		log.Fatal(err.Error())
	}

	isMatched, discoveredId, err := sdk.Identify(probe)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("isMatched : ", isMatched)
	fmt.Println("discovered Id: ", discoveredId)
}
