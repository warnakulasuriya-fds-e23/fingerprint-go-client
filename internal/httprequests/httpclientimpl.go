package httprequests

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/warnakulasuriya-fds-e23/fingerprint-go-client/internal/configtomlreader"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-sdk/core"
	"github.com/warnakulasuriya-fds-e23/go-sourceafis-fork/templates"
)

const (
	MatchTemplatesEndpoint    = "/api/fingerprint/match"
	IdentifyTemplateEndpoint  = "/api/fingerprint/identify"
	EnrollTemplateEndpoint    = "/api/fingerprint/enroll"
	AuthorizeTemplateEndpoint = "/api/fingerprint/authorize"
	UploadCborZipFileEndpoint = "/api/gallery/upload-cbor-zip"
)

type httpHeaderKeyValue struct {
	key   string
	value string
}

type Httpclientimpl struct {
	orchestrationServerAdress string
	imagesDir                 string
	cborDir                   string
	sdk                       *core.SDKCore
	headerKeyValueArray       []httpHeaderKeyValue
	accessToken               string
	expiryTime                time.Time
	mutex                     sync.Mutex
}

func NewHttpClientImpl() *Httpclientimpl {
	config := configtomlreader.ConfigTomlReader()

	sdk, err := core.NewSDKCore(config.ImagesDir, config.CborDir)
	if err != nil {
		log.Fatal(err.Error())
	}

	c := &Httpclientimpl{
		orchestrationServerAdress: os.Getenv("ORCHESTRATION_SERVER_ADRESS"),
		imagesDir:                 config.ImagesDir,
		cborDir:                   config.CborDir,
		sdk:                       sdk,
	}
	c.getAccessToken()
	return c
}

func (client *Httpclientimpl) SetOrAddHeaderValueAccordingToKey(key string, value string) {
	discoverkey := false
	for _, headerKeyValuePair := range client.headerKeyValueArray {
		if headerKeyValuePair.key == key {
			discoverkey = true
			headerKeyValuePair.value = value
			break
		}
	}
	if !discoverkey {
		client.headerKeyValueArray = append(client.headerKeyValueArray, httpHeaderKeyValue{key: key, value: value})
	} else {
		return
	}
}

func (client *Httpclientimpl) ClearAddedHeaderKeyValuePairs() {
	client.headerKeyValueArray = make([]httpHeaderKeyValue, 0)
}

func (client *Httpclientimpl) MatchTemplates(probe *templates.SearchTemplate, candidate *templates.SearchTemplate) (isMatch bool, err error) {

	isMatch, err = client.matchTemplateRequest(probe, candidate)
	return
}
func (client *Httpclientimpl) IdentifyTemplate(probe *templates.SearchTemplate) (isMatched bool, discoveredId string, err error) {
	isMatched, discoveredId, err = client.identifyTemplateRequest(probe)
	return
}
func (client *Httpclientimpl) EnrollTemplate(newEntry *templates.SearchTemplate, id string) (message string, err error) {
	message, err = client.enrollTemplateRequest(newEntry, id)
	return
}
func (client *Httpclientimpl) AuthorizeTemplate(template *templates.SearchTemplate) (status string, err error) {
	status, err = client.authorizeTemplateRequest(template)
	return
}
func (client *Httpclientimpl) MatchTemplatesFilesMethod(probeFilePath string, candidateFilePath string) (isMatch bool, err error) {
	// TODO: move main content of function body to a seperate file
	probe, err := client.sdk.Extract(probeFilePath)
	if err != nil {
		isMatch = false
		err = fmt.Errorf("error occured while extracting the template for the probe file, %w", err)
		return
	}
	candidate, err := client.sdk.Extract(candidateFilePath)
	if err != nil {
		isMatch = false
		err = fmt.Errorf("error occured while extractin the template for the candidate file, %w", err)
		return
	}
	isMatch, err = client.matchTemplateRequest(probe, candidate)
	return
}
func (client *Httpclientimpl) IdentifyTemplateFilesMethod(probeFilePath string) (isMatched bool, discoveredId string, err error) {
	// TODO: move main content of function body to a seperate file
	probe, err := client.sdk.Extract(probeFilePath)
	if err != nil {
		isMatched = false
		discoveredId = "none"
		err = fmt.Errorf("error occured while extracting the template for the probe file, %w", err)
		return
	}
	isMatched, discoveredId, err = client.identifyTemplateRequest(probe)
	return
}
func (client *Httpclientimpl) EnrollTemplateFilesMethod(newEntryFilePath string, id string) (message string, err error) {
	newEntry, err := client.sdk.Extract(newEntryFilePath)
	if err != nil {
		message = ""
		err = fmt.Errorf("error occured while extracting the template for the new Entry file, %w", err)
		return
	}
	message, err = client.enrollTemplateRequest(newEntry, id)
	return
}
func (client *Httpclientimpl) AuthorizeTemplateFilesMethod(templateFilePath string) (status string, err error) {
	status = "client side extracting template from file"
	template, err := client.sdk.Extract(templateFilePath)
	if err != nil {
		err = fmt.Errorf("error occured while extracting the template for the probe file, %w", err)
		return
	}
	status, err = client.authorizeTemplateRequest(template)
	return
}
func (client *Httpclientimpl) UploadCborZipFile(zipFilePath string) error {
	return client.uploadCborZipFile(zipFilePath)
}
