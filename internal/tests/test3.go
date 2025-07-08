package tests

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/warnakulasuriya-fds-e23/fingerprint-go-sdk/entities"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-sdk/sdkutils"
	"github.com/warnakulasuriya-fds-e23/go-sourceafis-fork"
)

func Test3() {
	fmt.Println("identify test using go sourceafis")
	var fingerprintImagesDirectoryPath string = "/home/dheera/FingerPrintDatabases/veryLargePNGDataset/"
	// var cborDirectoryPath = "/home/dheera/FingerPrintDatabases/cborDirectory2ForGo/"

	var gallery []*entities.SearchTemplateRecord
	sdkutils.LoadImagesDirToGallery(&gallery, fingerprintImagesDirectoryPath)

	probeImageName := "DB1_B105_7.png"
	probeImg, err := sourceafis.LoadImage(filepath.Join(fingerprintImagesDirectoryPath, probeImageName))
	if err != nil {
		log.Fatal(err.Error())
	}

	l := sourceafis.NewTransparencyLogger(new(TransparencyContents))
	tc := sourceafis.NewTemplateCreator(l)

	probe, err := tc.Template(probeImg)
	if err != nil {
		log.Fatal(err.Error())
	}

	matcher, err := sourceafis.NewMatcher(l, probe)
	if err != nil {
		log.Fatal(err.Error())
	}

	ctx := context.Background()
	max := -1000.000
	var match entities.SearchTemplateRecord
	threshold := 40
	t := time.Now()
	for _, templateRecordptr := range gallery {
		candidate := templateRecordptr.Template
		score := matcher.Match(ctx, &candidate)
		fmt.Println(score)
		if score >= max {
			max = score
			match = *templateRecordptr
		}
	}
	fmt.Println("max score: ", max)
	if max > float64(threshold) {
		fmt.Println("Match found : ", match.Id)
	} else {
		fmt.Println("No Match found")
	}
	fmt.Println("Time elapsed for 1:N match: ", time.Since(t))
}
