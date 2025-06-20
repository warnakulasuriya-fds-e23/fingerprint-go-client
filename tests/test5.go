package tests

import (
	"fmt"
	"log"

	"github.com/warnakulasuriya-fds-e23/fingerprint-go-sdk/core"
)

func Test5() {
	var fingerprintImagesDirectoryPath string = "/home/dheera/FingerPrintDatabases/smallPNGDataset/"
	var cborDirectoryPath = "/home/dheera/FingerPrintDatabases/cborDirectory2ForGo/"
	sdk, err := core.NewsdkCore(fingerprintImagesDirectoryPath, cborDirectoryPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	sdk.LoadImages()

	var probeImagePath string = "/home/dheera/FingerPrintDatabases/veryLargePNGDataset/DB3_B107_1.png"
	probe := sdk.Extract(probeImagePath)

	isMatched, discoveredId := sdk.Identify(probe)
	fmt.Println("isMatched : ", isMatched)
	fmt.Println("discovered Id: ", discoveredId)
}
