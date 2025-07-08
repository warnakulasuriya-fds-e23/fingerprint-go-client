package tests

import (
	"fmt"
	"log"

	httprequests "github.com/warnakulasuriya-fds-e23/fingerprint-go-client/internal/httprequests"
)

func Test10(client *httprequests.Httpclientimpl) {
	fmt.Println("cbor zip upload test <broken>")
	err := client.UploadCborZipFile("/home/dheera/FingerPrintDatabases/duplicateImpressionsFiltered/cborGoDirectory.zip")
	if err != nil {
		log.Fatal(err.Error())
	}
}
