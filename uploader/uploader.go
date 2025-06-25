package uploader

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/warnakulasuriya-fds-e23/biometric-orchestration-go-server/configuration"
	"github.com/warnakulasuriya-fds-e23/biometric-orchestration-go-server/responseobjects"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-client/requestobjects"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-sdk/core"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-sdk/entities"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-sdk/sdkutils"
)

//	type Uploader struct {
//		uploadTo string
//	}
func Uploader() {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}

	tomlpath := filepath.Join(workingDir, "config.toml")
	var config configuration.Configuration
	toml.DecodeFile(tomlpath, &config)

	sdk, err := core.NewSDKCore(config.ImagesDir, config.CborDir)
	if err != nil {
		log.Fatal(err.Error())
	}
	var galleryDummy []*entities.SearchTemplateRecord
	sdkutils.LoadCborDirToGallery(&galleryDummy, "/home/dheera/FingerPrintDatabases/cborDirectoryForGo/")
	for _, record := range galleryDummy {
		dataArray, err := sdk.GetAsByteArray(&record.Template)
		if err != nil {
			log.Fatal(err.Error())
		}
		reqObj := requestobjects.EnrollTemplateReqObj{Data: *dataArray, Id: record.Id}
		jsonobj, err := json.Marshal(reqObj)
		if err != nil {
			log.Fatal(err.Error())
		}

		requestBody := bytes.NewBuffer(jsonobj)
		fmt.Printf("sending request for id: %s\n", record.Id)
		req, err := http.NewRequest("POST", "https://3ee2e1ea-a3ec-4f5b-bff1-2fa3a86782c0-dev.e1-us-east-azure.choreoapis.dev/default/biometric-orchestration-s/v1.0/api/fingerprint/enroll", requestBody)
		if err != nil {
			log.Fatal(err.Error())
		}
		req.Header.Add("accept", "text/plain")
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Test-Key", "eyJraWQiOiJnYXRld2F5X2NlcnRpZmljYXRlX2FsaWFzIiwiYWxnIjoiUlMyNTYifQ.eyJzdWIiOiIxN2ZkZDg1ZS1hNWVjLTRhNzktOWEyMS0yY2QyNGI1NWJhY2NAY2FyYm9uLnN1cGVyIiwiYXVkIjoiY2hvcmVvOmRlcGxveW1lbnQ6c2FuZGJveCIsIm9yZ2FuaXphdGlvbiI6eyJ1dWlkIjoiM2VlMmUxZWEtYTNlYy00ZjViLWJmZjEtMmZhM2E4Njc4MmMwIn0sImlzcyI6Imh0dHBzOlwvXC9zdHMuY2hvcmVvLmRldjo0NDNcL2FwaVwvYW1cL3B1Ymxpc2hlclwvdjJcL2FwaXNcL2ludGVybmFsLWtleSIsImtleXR5cGUiOiJTQU5EQk9YIiwic3Vic2NyaWJlZEFQSXMiOlt7InN1YnNjcmliZXJUZW5hbnREb21haW4iOm51bGwsIm5hbWUiOiJiaW9tZXRyaWMtb3JjaGVzdHJhdGlvbi1zIC0gYmlvbWV0cmljLW9yY2hlc3RyYXRpb24tc2VydmVyIiwiY29udGV4dCI6IlwvM2VlMmUxZWEtYTNlYy00ZjViLWJmZjEtMmZhM2E4Njc4MmMwXC9kZWZhdWx0XC9iaW9tZXRyaWMtb3JjaGVzdHJhdGlvbi1zXC92MS4wIiwicHVibGlzaGVyIjoiY2hvcmVvX3Byb2RfYXBpbV9hZG1pbiIsInZlcnNpb24iOiJ2MS4wIiwic3Vic2NyaXB0aW9uVGllciI6bnVsbH1dLCJleHAiOjE3NTA4NDEwMzYsInRva2VuX3R5cGUiOiJJbnRlcm5hbEtleSIsImlhdCI6MTc1MDg0MDQzNiwianRpIjoiNWZlZjc1MmMtODUwNC00OTYwLWEwNmYtN2U2M2MxNmQ3NmE1In0.oHHoH3hCIjS1Wf26zXvWWqXZxGl-6zH9S3gjEuS6vw7uGN74H0YHcrHHsHn2QrKpYY_3G3Orw6QCDdjr3hdOFFgfuSAFYLUsGONg3uJoyHldPY37RPmixmOPHROFOQ7FgZuuKs0ZViX89B9i8JIvNGuPby5lw87c5FqSZ8MYyxzWDGAOauKhefmsvLH-ds-ak1hmDv3HzL7iONHdad5IvU1RqMuqZ1pIhqxJPENwyPmvW8EfJQJrXRndWKN1Zj7i6NxckbKOGpkdATqbb2vwH01etLdpazRSB9gpcmMeOsufG0wPOc6v9269lGTwwU3D7X0h1r4AwdL_h0UktJ7wQymmgqE5QmXGUfem7WssmH-1Ju2Png6UUN0jVX3f-aczl8DCwKEj0bZknWgA7_IwpLZTyHRmtOJ1dh0Lsk4qCacbrOIIvJV2DywLYjb-7_9FLQggpfx7H69s7Wp5uqs1_IXdFNdtM2F_rDdWQ-T1aVQuaet5W2V_Jk1bDUoG_sbcI9DihsZHEuAy6CKPjChzZ91kqq7h5UMVx8hyAhg_85VNgh3facfvjvZs7yqeNYtNWPeGtKkNfpdZ8Pje-Kpw_3KvsDLGE20v6XRpgeJnxXn9KW1RomN8jMiusXRFnZ5GsN4nJ1V1tR3pw2hMq7-vHTIXqvVoPUPxGRCCFxLoJKs")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer resp.Body.Close()
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error! %s\n", err.Error())
		}
		var resObj responseobjects.EnrollTemplateResObj
		err = json.Unmarshal(bodyBytes, &resObj)
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Printf("for the id %s, got response: %s\n", record.Id, resObj.Message)
	}

}
