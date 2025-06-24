package httprequests

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-client/configuration"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-sdk/core"
	"github.com/warnakulasuriya-fds-e23/go-sourceafis-fork/templates"
)

const (
	MatchTemplatesEndpoint   = "/api/fingerprint/match"
	IdentifyTemplateEndpoint = "/api/fingerprint/identify"
	EnrollTemplateEndpoint   = "/api/fingerprint/enroll"
)

type httpclientimpl struct {
	orchestrationServerAdress string
	imagesDir                 string
	cborDir                   string
	sdk                       *core.SDKCore
}

func NewHttpClientImpl() *httpclientimpl {
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

	c := &httpclientimpl{
		orchestrationServerAdress: config.OrchestrationServerAdress,
		imagesDir:                 config.ImagesDir,
		cborDir:                   config.CborDir,
		sdk:                       sdk,
	}
	return c
}

// TODO: Implement proper error handling for the http Request methods
func (client *httpclientimpl) MatchTemplates(probe *templates.SearchTemplate, candidate *templates.SearchTemplate) (isMatch bool) {

	isMatch = client.matchTemplate(probe, candidate)
	return
}
func (client *httpclientimpl) IdentifyTemplate(probe *templates.SearchTemplate) (isMatched bool, discoveredId string) {
	isMatched, discoveredId = client.identifyTemplateRequest(probe)
	return
}
func (client *httpclientimpl) EnrollTemplate(newEntry *templates.SearchTemplate, id string) (message string, err error) {
	message, err = client.enrollTemplateRequest(newEntry, id)
	return
}
func (client *httpclientimpl) MatchTemplatesFilesMethod(probeFilePath string, candidateFilePath string) (isMatch bool, err error) {
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
	isMatch = client.matchTemplate(probe, candidate)
	err = nil
	return
}
func (client *httpclientimpl) IdentifyTemplateFilesMethod(probeFilePath string) (isMatched bool, discoveredId string, err error) {
	// TODO: move main content of function body to a seperate file
	probe, err := client.sdk.Extract(probeFilePath)
	if err != nil {
		isMatched = false
		discoveredId = "none"
		err = fmt.Errorf("error occured while extracting the template for the probe file, %w", err)
		return
	}
	isMatched, discoveredId = client.identifyTemplateRequest(probe)
	err = nil
	return
}
func (client *httpclientimpl) EnrollTemplateFilesMethod(newEntryFilePath string, id string) (message string, err error) {
	newEntry, err := client.sdk.Extract(newEntryFilePath)
	if err != nil {
		message = ""
		err = fmt.Errorf("error occured while extracting the template for the new Entry file, %w", err)
		return
	}
	message, err = client.enrollTemplateRequest(newEntry, id)
	return
}
