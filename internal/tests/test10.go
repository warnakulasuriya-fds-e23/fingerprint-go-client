package tests

import (
	"fmt"
	"log"

	"github.com/warnakulasuriya-fds-e23/fingerprint-go-client/internal/outgoing"
)

func Test10(client *outgoing.OutgoingClient) {
	fmt.Println("cbor zip upload test <broken>")
	err := client.UploadCborZipFile("/home/dheera/FingerPrintDatabases/duplicateImpressionsFiltered/cborGoDirectory.zip")
	if err != nil {
		log.Fatal(err.Error())
	}
}
