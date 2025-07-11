package httprequests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/warnakulasuriya-fds-e23/fingerprint-go-client/internal/requestobjects"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-client/internal/responseobjects"
	"github.com/warnakulasuriya-fds-e23/go-sourceafis-fork/templates"
)

func (client *Httpclientimpl) identifyTemplateRequest(probe *templates.SearchTemplate) (isMatched bool, discoveredId string, err error) {
	var accessToken string
	if client.accessToken == "" || client.expiryTime.Equal(time.Now()) || client.expiryTime.Before(time.Now().Add(5*time.Second)) {
		accessToken, err = client.getAccessToken()
		if err != nil {
			err = fmt.Errorf("identifyTemplateRequest] error while trying get access token ,either new token or existing token, %w", err)
			return
		}
	} else {
		accessToken = client.accessToken
	}
	client.SetOrAddHeaderValueAccordingToKey("Content-Type", "application/json")
	probeBytes, err := client.sdk.GetAsByteArray(probe)
	if err != nil {
		isMatched = false
		discoveredId = "none"
		err = fmt.Errorf("[identifyTemplateRequest] error while running GetAsByteArray for probe template : %w", err)
		return
	}
	reqobj := requestobjects.IdentifyTemplateReqObj{ProbeCbor: *probeBytes, DeviceId: os.Getenv("DEVICE_ID")}
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

	if resp.StatusCode != 200 {
		var resObj responseobjects.ErrorResObj
		err = json.Unmarshal(bodyBytes, &resObj)
		if err != nil {
			err = fmt.Errorf("[identifyTemplateRequest] error occured while runnig json.Unmarshal on response bytes , %w", err)
			return
		}
		isMatched = false
		discoveredId = "none"
		err = fmt.Errorf("[identifyTemplateRequest] error occured in sending to orchestration service , %s", resObj.Message)
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
