package tests

import (
	"fmt"
	"log"

	httprequests "github.com/warnakulasuriya-fds-e23/fingerprint-go-client/httpRequests"
)

func Test8() {
	client := httprequests.NewHttpClientImpl()
	var NewEntryImagePath string = "/home/dheera/FingerPrintDatabases/veryLargePNGDataset/DB3_B107_1.png"
	message, err := client.EnrollTemplateFilesMethod(NewEntryImagePath, "testuser1")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(message)
}
