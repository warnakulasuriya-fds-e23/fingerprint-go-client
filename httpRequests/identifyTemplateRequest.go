package httprequests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/warnakulasuriya-fds-e23/fingerprint-go-client/requestobjects"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-client/responseobjects"
	"github.com/warnakulasuriya-fds-e23/go-sourceafis-fork/templates"
)

func (client *Httpclientimpl) identifyTemplateRequest(probe *templates.SearchTemplate) (isMatched bool, discoveredId string, err error) {
	accessToken := client.getAccessToken()
	client.SetOrAddHeaderValueAccordingToKey("Content-Type", "application/json")
	probeBytes, err := client.sdk.GetAsByteArray(probe)
	if err != nil {
		isMatched = false
		discoveredId = "none"
		err = fmt.Errorf("[identifyTemplateRequest] error while running GetAsByteArray for probe template : %w", err)
		return
	}
	reqobj := requestobjects.IdentifyTemplateReqObj{ProbeCbor: *probeBytes}
	jsonobj, err := json.Marshal(reqobj)
	if err != nil {
		isMatched = false
		discoveredId = "none"
		err = fmt.Errorf("[identifyTemplateRequest] error while running json marshal for identify template request object : %w", err)
		return
	}
	urlString, err := url.JoinPath(client.orchestrationServerAdress, IdentifyTemplateEndpoint)
	if err != nil {
		isMatched = false
		discoveredId = "none"
		err = fmt.Errorf("[identifyTemplateRequest] error while trying join the base url %s with %s : %w", client.orchestrationServerAdress, IdentifyTemplateEndpoint, err)
		return
	}
	requestBody := bytes.NewBuffer(jsonobj)
	req, err := http.NewRequest("POST", urlString, requestBody)
	if err != nil {
		isMatched = false
		discoveredId = "none"
		err = fmt.Errorf("[identifyTemplateRequest] error while creating a new post request : %w", err)
		return
	}

	client.SetOrAddHeaderValueAccordingToKey("Authorization", "Bearer "+accessToken)
	for _, headerKeyValuePair := range client.headerKeyValueArray {
		req.Header.Add(headerKeyValuePair.key, headerKeyValuePair.value)
	}

	internalClient := &http.Client{}
	resp, err := internalClient.Do(req)
	if err != nil {
		isMatched = false
		discoveredId = "none"
		err = fmt.Errorf("[identifyTemplateRequest] error while sending identify post request or getting response: %w", err)
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		isMatched = false
		discoveredId = "none"
		err = fmt.Errorf("[identifyTemplateRequest] error while trying to read bytest from response body: %w", err)
		return
	}
	var resobj responseobjects.IdentifyTemplateResObje
	err = json.Unmarshal(bodyBytes, &resobj)
	if err != nil {
		isMatched = false
		discoveredId = "none"
		err = fmt.Errorf("[identifyTemplateRequest] error while trying to json unmarshal the bytes read from request body: %w", err)
		return
	}
	isMatched = resobj.IsMatched
	discoveredId = resobj.DiscoveredId
	return
}
