package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fxamacker/cbor/v2"
	"github.com/miqdadyyy/go-sourceafis/templates"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-client/entities"
)

func loadCborfileToGallery(galleryptr *[]*entities.SearchTemplateRecord, cborpath string, filename string) {
	readData, err := os.ReadFile(cborpath)
	if err != nil {
		fmt.Println("Error in reading the file at", cborpath)
		return
	}
	var template templates.SearchTemplate
	err = cbor.Unmarshal(readData, &template)
	if err != nil {
		fmt.Println("Error when while executing unmarshal function for data from ", cborpath)
	}
	fileNameWithoutExtenstion := strings.TrimSuffix(filename, filepath.Ext(cborpath))
	*galleryptr = append(*galleryptr, &entities.SearchTemplateRecord{Id: fileNameWithoutExtenstion, Template: template})
	fmt.Printf("appended template from %s to gallery\n", filename)
}

func LoadCborDirToGallery(galleryptr *[]*entities.SearchTemplateRecord, cborDir string) {
	files, err := os.ReadDir(cborDir)
	if err != nil {
		fmt.Println("Error! Unable to read data")
	}

	for _, file := range files {
		loadCborfileToGallery(galleryptr, filepath.Join(cborDir, file.Name()), file.Name())
	}

}
