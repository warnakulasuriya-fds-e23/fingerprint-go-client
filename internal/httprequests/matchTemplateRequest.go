package httprequests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/warnakulasuriya-fds-e23/fingerprint-go-client/internal/requestobjects"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-client/internal/responseobjects"
	"github.com/warnakulasuriya-fds-e23/go-sourceafis-fork/templates"
)

func (client *Httpclientimpl) matchTemplateRequest(probe *templates.SearchTemplate, candidate *templates.SearchTemplate) (isMatch bool, err error) {
	var accessToken string
	if client.accessToken == "" || client.expiryTime.Equal(time.Now()) || client.expiryTime.Before(time.Now().Add(5*time.Second)) {
		accessToken, err = client.getAccessToken()
		if err != nil {
			err = fmt.Errorf("error while trying get access token ,either new token or existing token, %w", err)
			return
		}
	} else {
		accessToken = client.accessToken
	}
	client.SetOrAddHeaderValueAccordingToKey("Content-Type", "application/json")
	probeBytes, err := client.sdk.GetAsByteArray(probe)
	if err != nil {
		isMatch = false
		err = fmt.Errorf("[matchTemplateRequest] error while trying to run get as byte array method for probe template: %w", err)
		return
	}
	candidateBytes, err := client.sdk.GetAsByteArray(candidate)
	if err != nil {
		isMatch = false
		err = fmt.Errorf("[matchTemplateRequest] error while trying to run get as byte array method for candidate template: %w", err)
		return
	}

	obj := requestobjects.MatchTemplatesReqObj{ProbeCbor: *probeBytes, CandidateCbor: *candidateBytes}
	jsonobj, err := json.Marshal(obj)
	if err != nil {
		isMatch = false
		err = fmt.Errorf("[matchTemplateRequest] error while trying json marhsal match template request object: %w", err)
		return
	}
	urlString, err := url.JoinPath(client.orchestrationServerAdress, MatchTemplatesEndpoint)
	if err != nil {
		isMatch = false
		err = fmt.Errorf("[matchTemplateRequest] error while trying to combine url %s with %s : %w", client.orchestrationServerAdress, MatchTemplatesEndpoint, err)
		return
	}
	requestBody := bytes.NewBuffer(jsonobj)
	req, err := http.NewRequest("POST", urlString, requestBody)
	if err != nil {
		isMatch = false
		err = fmt.Errorf("[matchTemplateRequest] error while making new post request for match template: %w", err)
		return
	}

	client.SetOrAddHeaderValueAccordingToKey("Authorization", "Bearer "+accessToken)
	for _, headerKeyValuePair := range client.headerKeyValueArray {
		req.Header.Add(headerKeyValuePair.key, headerKeyValuePair.value)
	}

	internalClient := &http.Client{}
	resp, err := internalClient.Do(req)
	if err != nil {
		isMatch = false
		err = fmt.Errorf("[matchTemplateRequest] error while sending match template http request or recieving response : %w", err)
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		isMatch = false
		err = fmt.Errorf("[matchTemplateRequest] error while reading bytes of response body : %w", err)
		return
	}
	var responseobj responseobjects.MatchTemplatesResObj
	err = json.Unmarshal(bodyBytes, &responseobj)
	if err != nil {
		isMatch = false
		err = fmt.Errorf("[matchTemplateRequest] error while trying to json unmarshal the bytes read from the request body : %w", err)
		return
	}
	isMatch = responseobj.IsMatch
	return
}
