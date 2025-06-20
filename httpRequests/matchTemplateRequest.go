package httprequests

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/warnakulasuriya-fds-e23/fingerprint-go-client/requestobjects"
	"github.com/warnakulasuriya-fds-e23/go-sourceafis-fork/templates"
)

func (client *httpclientimpl) matchTemplate(probe templates.SearchTemplate, candidate templates.SearchTemplate) (isMatch bool) {
	probeBytes := client.sdk.GetAsByteArray(&probe)
	candidateBytes := client.sdk.GetAsByteArray(&candidate)

	obj := requestobjects.MatchTemplatesReqObj{ProbeCbor: *probeBytes, CandidateCbor: *candidateBytes}
	jsonobj, err := json.Marshal(obj)
	if err != nil {
		log.Fatal(err.Error())
	}
	url, err := url.JoinPath(client.orchestrationServerAdress, MatchTemplatesEndpoint)
	if err != nil {
		log.Fatal(err.Error())
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonobj))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	bdy := string(body)

	if bdy == "match" {
		return true
	} else {
		return false
	}
}
