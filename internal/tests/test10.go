package tests

import (
	"log"

	httprequests "github.com/warnakulasuriya-fds-e23/fingerprint-go-client/internal/httprequests"
)

func Test10(client *httprequests.Httpclientimpl) {
	err := client.UploadCborZipFile("/home/dheera/FingerPrintDatabases/duplicateImpressionsFiltered/cborGoDirectory.zip")
	if err != nil {
		log.Fatal(err.Error())
	}
}
