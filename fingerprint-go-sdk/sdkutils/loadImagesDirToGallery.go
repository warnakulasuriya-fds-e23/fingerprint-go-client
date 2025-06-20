package sdkutils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/miqdadyyy/go-sourceafis"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-sdk/entities"
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

func loadImageToGallery(galleryptr *[]*entities.SearchTemplateRecord, imagePath string, filename string) {
	image, err := sourceafis.LoadImage(imagePath)
	if err != nil {
		fmt.Println("Couldnt LoadImage at path: ", imagePath)
		return
	}
	l := sourceafis.NewTransparencyLogger(new(TransparencyContents))
	tc := sourceafis.NewTemplateCreator(l)

	template, err := tc.Template(image)
	if err != nil {
		fmt.Println("Couldnt extract template from Image at path: ", imagePath)
		return
	}
	fileNameWithoutExtenstion := strings.TrimSuffix(filename, filepath.Ext(filename))
	*galleryptr = append(*galleryptr, &entities.SearchTemplateRecord{Id: fileNameWithoutExtenstion, Template: *template})

	fmt.Printf("appended template of %s to gallery\n", filename)

}

func LoadImagesDirToGallery(galleryptr *[]*entities.SearchTemplateRecord, imagesDir string) {
	// iterate over the files in the imagesDir get their absolute path and
	// for every iteration execute the loadImageToGallery function by passing that path
	// as an argument
	files, err := os.ReadDir(imagesDir)
	if err != nil {
		fmt.Println("error reading directory")
		return
	}

	for _, file := range files {
		filename := file.Name()
		loadImageToGallery(galleryptr, filepath.Join(imagesDir, filename), filename)
	}
}
