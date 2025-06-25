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

func (client *Httpclientimpl) matchTemplate(probe *templates.SearchTemplate, candidate *templates.SearchTemplate) (isMatch bool) {
	client.SetOrAddHeaderValueAccordingToKey("Content-Type", "application/json")
	probeBytes, err := client.sdk.GetAsByteArray(probe)
	if err != nil {
		log.Fatal(err.Error())
	}
	candidateBytes, err := client.sdk.GetAsByteArray(candidate)
	if err != nil {
		log.Fatal(err.Error())
	}

	obj := requestobjects.MatchTemplatesReqObj{ProbeCbor: *probeBytes, CandidateCbor: *candidateBytes}
	jsonobj, err := json.Marshal(obj)
	if err != nil {
		log.Fatal(err.Error())
	}
	urlString, err := url.JoinPath(client.orchestrationServerAdress, MatchTemplatesEndpoint)
	if err != nil {
		log.Fatal(err.Error())
	}
	requestBody := bytes.NewBuffer(jsonobj)
	req, err := http.NewRequest("POST", urlString, requestBody)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, headerKeyValuePair := range client.headerKeyValueArray {
		req.Header.Add(headerKeyValuePair.key, headerKeyValuePair.value)
	}

	internalClient := &http.Client{}
	resp, err := internalClient.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	var responseobj responseobjects.MatchTemplatesResObj
	err = json.Unmarshal(bodyBytes, &responseobj)
	if err != nil {
		log.Fatal(err.Error())
	}
	isMatch = responseobj.IsMatch
	return
}
