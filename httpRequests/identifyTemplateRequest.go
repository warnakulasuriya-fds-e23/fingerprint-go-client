package httprequests

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/warnakulasuriya-fds-e23/fingerprint-go-client/requestobjects"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-client/responseobjects"
	"github.com/warnakulasuriya-fds-e23/go-sourceafis-fork/templates"
)

func (client *httpclientimpl) identifyTemplateRequest(probe *templates.SearchTemplate) (isMatched bool, discoveredId string) {
	probeBytes, err := client.sdk.GetAsByteArray(probe)
	if err != nil {
		log.Fatal(err.Error())
	}
	reqobj := requestobjects.IdentifyTemplateReqObj{ProbeCbor: *probeBytes}
	jsonobj, err := json.Marshal(reqobj)
	if err != nil {
		log.Fatal(err.Error())
	}
	urlString, err := url.JoinPath(client.orchestrationServerAdress, IdentifyTemplateEndpoint)
	if err != nil {
		log.Fatal(err.Error())
	}

	resp, err := http.Post(urlString, "application/json", bytes.NewBuffer(jsonobj))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	var resobj responseobjects.IdentifyTemplateResObje
	err = json.Unmarshal(bodyBytes, &resobj)
	if err != nil {
		log.Fatal(err.Error())
	}
	isMatched = resobj.IsMatched
	discoveredId = resobj.DiscoveredId
	return
}
