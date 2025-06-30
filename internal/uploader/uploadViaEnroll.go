package uploader

import (
	"fmt"

	httprequests "github.com/warnakulasuriya-fds-e23/fingerprint-go-client/httpRequests"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-sdk/entities"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-sdk/sdkutils"
)

//	type Uploader struct {
//		uploadTo string
//	}
func UploadViaEnroll(client *httprequests.Httpclientimpl) {
	var galleryDummy []*entities.SearchTemplateRecord
	sdkutils.LoadCborDirToGallery(&galleryDummy, "/home/dheera/FingerPrintDatabases/duplicateImpressionsFiltered/cborGoDirectory/")
	for _, record := range galleryDummy {

		// req.Header.Add("accept", "text/plain")
		// req.Header.Add("Content-Type", "application/json")
		// req.Header.Add("Test-Key", "eyJraWQiOiJnYXRld2F5X2NlcnRpZmljYXRlX2FsaWFzIiwiYWxnIjoiUlMyNTYifQ.eyJzdWIiOiIxN2ZkZDg1ZS1hNWVjLTRhNzktOWEyMS0yY2QyNGI1NWJhY2NAY2FyYm9uLnN1cGVyIiwiYXVkIjoiY2hvcmVvOmRlcGxveW1lbnQ6c2FuZGJveCIsIm9yZ2FuaXphdGlvbiI6eyJ1dWlkIjoiM2VlMmUxZWEtYTNlYy00ZjViLWJmZjEtMmZhM2E4Njc4MmMwIn0sImlzcyI6Imh0dHBzOlwvXC9zdHMuY2hvcmVvLmRldjo0NDNcL2FwaVwvYW1cL3B1Ymxpc2hlclwvdjJcL2FwaXNcL2ludGVybmFsLWtleSIsImtleXR5cGUiOiJTQU5EQk9YIiwic3Vic2NyaWJlZEFQSXMiOlt7InN1YnNjcmliZXJUZW5hbnREb21haW4iOm51bGwsIm5hbWUiOiJiaW9tZXRyaWMtb3JjaGVzdHJhdGlvbi1zIC0gYmlvbWV0cmljLW9yY2hlc3RyYXRpb24tc2VydmVyIiwiY29udGV4dCI6IlwvM2VlMmUxZWEtYTNlYy00ZjViLWJmZjEtMmZhM2E4Njc4MmMwXC9kZWZhdWx0XC9iaW9tZXRyaWMtb3JjaGVzdHJhdGlvbi1zXC92MS4wIiwicHVibGlzaGVyIjoiY2hvcmVvX3Byb2RfYXBpbV9hZG1pbiIsInZlcnNpb24iOiJ2MS4wIiwic3Vic2NyaXB0aW9uVGllciI6bnVsbH1dLCJleHAiOjE3NTA4NDg3MTMsInRva2VuX3R5cGUiOiJJbnRlcm5hbEtleSIsImlhdCI6MTc1MDg0ODExMywianRpIjoiMTc2MTM0YzQtZTEyYi00ZTkzLTliMjAtODZjZjA1MmQ4YTRlIn0.Dy5c4TSzY1Azu_7MNw1cuJ9zzNKbgSncUsKEbe1mBb1odtuO_609heMuAkRQjIgGDNC5opGsIxCg_hRDdEDcYqN_Bbkow25UspaP4CLuUeCEWuaoHV1l9HQ98EKaJkrG1RP-pyLfy4r2quldgDdm0cu4KEbg9KSVVdMI8P5QEr_kAmmBYFxTKhCSEkeOg4x0YDNmB0wtfO8TMKrCvpbR5RWuH88N2riYHwbk6DD-twg4GHrVbq9BHl4i_4kBuZZeAXVvWJ9q1Tf38PUB28lVa80ZNt-SUmAKzxnOHbHhrG16-cOLdvrhUm8zve-sirgPhfqx_1C-kONQ8-8lDNgi8F-f9XGoItbhfBklK5SpQobtDkelsRoW0GkB2nlnJT5LbjQOcaa5TH4Pg--b7NCgSSg6zN7btysudipvjkdv912Y9_SxvElAHyYhvurpp9LRR1nb0Rp30smVM__GpSFGAsAZY7O8CLLRCfRrfGTUn27Dp9qRRMQXfCLtvDVlLVPtUnbSjbQybCxtOuS-FR8yKedqv0TCakdnNc4luw42L7VO70imYewvo0vwJh9YoTHlbqEekSCIvQPNidG1q-xtlwSPNtrddm0iNqArZxTcH0AZt0EwSb3QVC_IIt6JlMMW7hfoL-cKjQH6bVEO0JvPGit7RsYadHBM-Dk__QktApM")

		message, err := client.EnrollTemplate(&record.Template, record.Id)
		if err != nil {
			fmt.Printf("for the id %s, got error response: %s\n", record.Id, err.Error())
		} else {

			fmt.Printf("for the id %s, got response: %s\n", record.Id, message)
		}
	}

}
