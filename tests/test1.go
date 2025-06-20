package tests

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/warnakulasuriya-fds-e23/go-sourceafis-fork"
)

type TransparencyContents struct {
}

func (c *TransparencyContents) Accepts(key string) bool {
	return true
}

func (c *TransparencyContents) Accept(key, mime string, data []byte) error {
	//fmt.Printf("%d B  %s %s \n", len(data), mime, key)
	return nil
}

func Test1() {
	var probeImagePath string = "/home/dheera/FingerPrintDatabases/veryLargePNGDataset/DB1_B101_1.png"
	var matchingImagePath string = probeImagePath
	var nonMatchingImagePath string = "/home/dheera/FingerPrintDatabases/veryLargePNGDataset/DB1_B102_2.png"

	fmt.Println("testing image loading, template extraction and 1:1 matching")
	t := time.Now()
	probeImg, err := sourceafis.LoadImage(probeImagePath)
	fmt.Println("Time elapsed to Load probe Image: ", time.Since(t))

	if err != nil {
		log.Fatal(err.Error())
	}
	l := sourceafis.NewTransparencyLogger(new(TransparencyContents))
	tc := sourceafis.NewTemplateCreator(l)

	t = time.Now()
	probe, err := tc.Template(probeImg)
	fmt.Println("Time extract  probe Template using template creator: ", time.Since(t))
	if err != nil {
		log.Fatal(err.Error())
	}

	t = time.Now()
	candidateImg1, err := sourceafis.LoadImage(matchingImagePath)
	fmt.Println("Time elapsed to Load matching candidate Image: ", time.Since(t))
	if err != nil {
		log.Fatal(err.Error())
	}

	t = time.Now()
	candidate1, err := tc.Template(candidateImg1)
	fmt.Println("Time elapsed to extract template of matching candidate Image: ", time.Since(t))
	if err != nil {
		log.Fatal(err.Error())
	}

	t = time.Now()
	candidateImg2, err := sourceafis.LoadImage(nonMatchingImagePath)
	fmt.Println("Time elapsed to Load non matching candidate Image: ", time.Since(t))
	if err != nil {
		log.Fatal(err.Error())
	}

	t = time.Now()
	candidate2, err := tc.Template(candidateImg2)
	fmt.Println("Time elapsed to extract template from non matching candidate Image: ", time.Since(t))
	if err != nil {
		log.Fatal(err.Error())
	}

	matcher, err := sourceafis.NewMatcher(l, probe)
	if err != nil {
		log.Fatal(err.Error())
	}

	ctx := context.Background()

	t = time.Now()
	score1 := matcher.Match(ctx, candidate1)
	fmt.Println("Time elapsed to run matching process for probe and matching candidate: ", time.Since(t))

	t = time.Now()
	score2 := matcher.Match(ctx, candidate2)
	fmt.Println("Time elapsed to run matching process for probe and non matching candidate: ", time.Since(t))

	fmt.Printf("Score1 : %f \nScore2 : %f", score1, score2)

	// cborform, err := cbor.Marshal(probe)

	// fmt.Printf("CBOR form : ", cborform)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// var test templates.SearchTemplate

	// err = cbor.Unmarshal(cborform, &test)

	// t = time.Now()

}
