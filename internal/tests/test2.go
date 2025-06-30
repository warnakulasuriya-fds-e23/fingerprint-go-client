package tests

import (
	"fmt"
	"time"

	"github.com/warnakulasuriya-fds-e23/fingerprint-go-sdk/entities"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-sdk/sdkutils"
)

func Test2() {
	fmt.Println("loading files up from Image Dir and saving to Cbor Dir")
	var fingerprintImagesDirectoryPath string = "/home/dheera/FingerPrintDatabases/veryLargePNGDataset/"
	var cborDirectoryPath = "/home/dheera/FingerPrintDatabases/cborDirectory2ForGo/"

	var gallery []*entities.SearchTemplateRecord
	t := time.Now()
	sdkutils.LoadImagesDirToGallery(&gallery, fingerprintImagesDirectoryPath)
	firstduration := time.Since(t)
	fmt.Println("Time elapsed to load Images Directory to gallery: ", firstduration)

	t = time.Now()
	sdkutils.SaveGalleryToCborDir(&gallery, cborDirectoryPath)
	secondduration := time.Since(t)
	fmt.Println("Time elapsed to save gallery to cbor directory: ", secondduration)

	fmt.Println("clearing gallery")
	gallery = make([]*entities.SearchTemplateRecord, 0)
	fmt.Println("gallery cleared")

	t = time.Now()
	sdkutils.LoadCborDirToGallery(&gallery, cborDirectoryPath)
	thirdduration := time.Since(t)
	fmt.Println("Time elapsed to load cbor dir to gallery: ", thirdduration)
	fmt.Println("\n\n\nloading up images dir duration: ", firstduration)
	fmt.Println("saving gallery to cbor dir duration: ", secondduration)
	fmt.Println("loading cbor dir to clean gallery duration: ", thirdduration)

}
