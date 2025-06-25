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
	sdkutils.LoadCborDirToGallery(&galleryDummy, "/home/dheera/FingerPrintDatabases/duplicateImpressionsFiltered/cborGoDirectory/")
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
		// req, err := http.NewRequest("POST", "http://localhost:4000/api/fingerprint/enroll", requestBody)
		if err != nil {
			log.Fatal(err.Error())
		}
		req.Header.Add("accept", "text/plain")
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Test-Key", "eyJraWQiOiJnYXRld2F5X2NlcnRpZmljYXRlX2FsaWFzIiwiYWxnIjoiUlMyNTYifQ.eyJzdWIiOiIxN2ZkZDg1ZS1hNWVjLTRhNzktOWEyMS0yY2QyNGI1NWJhY2NAY2FyYm9uLnN1cGVyIiwiYXVkIjoiY2hvcmVvOmRlcGxveW1lbnQ6c2FuZGJveCIsIm9yZ2FuaXphdGlvbiI6eyJ1dWlkIjoiM2VlMmUxZWEtYTNlYy00ZjViLWJmZjEtMmZhM2E4Njc4MmMwIn0sImlzcyI6Imh0dHBzOlwvXC9zdHMuY2hvcmVvLmRldjo0NDNcL2FwaVwvYW1cL3B1Ymxpc2hlclwvdjJcL2FwaXNcL2ludGVybmFsLWtleSIsImtleXR5cGUiOiJTQU5EQk9YIiwic3Vic2NyaWJlZEFQSXMiOlt7InN1YnNjcmliZXJUZW5hbnREb21haW4iOm51bGwsIm5hbWUiOiJiaW9tZXRyaWMtb3JjaGVzdHJhdGlvbi1zIC0gYmlvbWV0cmljLW9yY2hlc3RyYXRpb24tc2VydmVyIiwiY29udGV4dCI6IlwvM2VlMmUxZWEtYTNlYy00ZjViLWJmZjEtMmZhM2E4Njc4MmMwXC9kZWZhdWx0XC9iaW9tZXRyaWMtb3JjaGVzdHJhdGlvbi1zXC92MS4wIiwicHVibGlzaGVyIjoiY2hvcmVvX3Byb2RfYXBpbV9hZG1pbiIsInZlcnNpb24iOiJ2MS4wIiwic3Vic2NyaXB0aW9uVGllciI6bnVsbH1dLCJleHAiOjE3NTA4NDg3MTMsInRva2VuX3R5cGUiOiJJbnRlcm5hbEtleSIsImlhdCI6MTc1MDg0ODExMywianRpIjoiMTc2MTM0YzQtZTEyYi00ZTkzLTliMjAtODZjZjA1MmQ4YTRlIn0.Dy5c4TSzY1Azu_7MNw1cuJ9zzNKbgSncUsKEbe1mBb1odtuO_609heMuAkRQjIgGDNC5opGsIxCg_hRDdEDcYqN_Bbkow25UspaP4CLuUeCEWuaoHV1l9HQ98EKaJkrG1RP-pyLfy4r2quldgDdm0cu4KEbg9KSVVdMI8P5QEr_kAmmBYFxTKhCSEkeOg4x0YDNmB0wtfO8TMKrCvpbR5RWuH88N2riYHwbk6DD-twg4GHrVbq9BHl4i_4kBuZZeAXVvWJ9q1Tf38PUB28lVa80ZNt-SUmAKzxnOHbHhrG16-cOLdvrhUm8zve-sirgPhfqx_1C-kONQ8-8lDNgi8F-f9XGoItbhfBklK5SpQobtDkelsRoW0GkB2nlnJT5LbjQOcaa5TH4Pg--b7NCgSSg6zN7btysudipvjkdv912Y9_SxvElAHyYhvurpp9LRR1nb0Rp30smVM__GpSFGAsAZY7O8CLLRCfRrfGTUn27Dp9qRRMQXfCLtvDVlLVPtUnbSjbQybCxtOuS-FR8yKedqv0TCakdnNc4luw42L7VO70imYewvo0vwJh9YoTHlbqEekSCIvQPNidG1q-xtlwSPNtrddm0iNqArZxTcH0AZt0EwSb3QVC_IIt6JlMMW7hfoL-cKjQH6bVEO0JvPGit7RsYadHBM-Dk__QktApM")

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
