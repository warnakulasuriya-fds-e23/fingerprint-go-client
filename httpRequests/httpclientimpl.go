package httprequests

import (
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-client/configuration"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-sdk/core"
	"github.com/warnakulasuriya-fds-e23/go-sourceafis-fork/templates"
)

const (
	MatchTemplatesEndpoint = "/api/fingerprint/match"
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

func (client *httpclientimpl) MatchTemplates(probe *templates.SearchTemplate, candidate *templates.SearchTemplate) (isMatch bool) {

	isMatch = client.matchTemplate(probe, candidate)
	return
}
