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

func (client *Httpclientimpl) authorizeTemplateRequest(template *templates.SearchTemplate) (status string, err error) {
	var accessToken string
	status = "client side Access Token setup"
	if client.accessToken == "" || client.expiryTime.Equal(time.Now()) || client.expiryTime.Before(time.Now().Add(5*time.Second)) {
		accessToken, err = client.getAccessToken()
		if err != nil {
			err = fmt.Errorf("error while trying get access token ,either new token or existing token, %w", err)
			return
		}
	} else {
		accessToken = client.accessToken
	}
	status = "client Side cbor conversion"
	probeBytes, err := client.sdk.GetAsByteArray(template)
	if err != nil {
		err = fmt.Errorf("[authorizeTemplateRequest] error while running GetAsByteArray for probe template : %w", err)
		return
	}

	status = "client Side preparing request body"
	reqobj := requestobjects.IdentifyTemplateReqObj{ProbeCbor: *probeBytes, DeviceId: os.Getenv("DEVICE_ID")}
	jsonobj, err := json.Marshal(reqobj)
	if err != nil {
		err = fmt.Errorf("[authorizeTemplateRequest] error while running json marshal for identify template request object : %w", err)
		return
	}

	requestBody := bytes.NewBuffer(jsonobj)

	status = "client Side preparing request url"
	urlString, err := url.JoinPath(client.orchestrationServerAdress, AuthorizeTemplateEndpoint)
	if err != nil {
		err = fmt.Errorf("[authorizeTemplateRequest] error while trying join the base url %s with %s : %w", client.orchestrationServerAdress, IdentifyTemplateEndpoint, err)
		return
	}

	status = "client Side preparing request"
	req, err := http.NewRequest("POST", urlString, requestBody)
	if err != nil {
		err = fmt.Errorf("[authorizeTemplateRequest] error while creating a new post request : %w", err)
		return
	}

	status = "client Side adding headers"
	// TODO: use net/http req.Header.Add instead of SetOrAddHeaderValueAccordingToKey
	client.SetOrAddHeaderValueAccordingToKey("Content-Type", "application/json")
	client.SetOrAddHeaderValueAccordingToKey("Authorization", "Bearer "+accessToken)
	for _, headerKeyValuePair := range client.headerKeyValueArray {
		req.Header.Add(headerKeyValuePair.key, headerKeyValuePair.value)
	}

	internalClient := &http.Client{}

	status = "client Side sending req"
	resp, err := internalClient.Do(req)
	if err != nil {
		err = fmt.Errorf("[authorizeTemplateRequest] error while sending identify post request or getting response: %w", err)
		return
	}
	defer resp.Body.Close()

	status = "client Side reading response body"
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("[authorizeTemplateRequest] error while trying to read bytest from response body: %w", err)
		return
	}

	if resp.StatusCode != 200 {
		var resObj responseobjects.ErrorResObj
		status = "client Side unmarshal error response body"
		err = json.Unmarshal(bodyBytes, &resObj)
		if err != nil {
			err = fmt.Errorf("[authorizeTemplateRequest] error occured while runnig json.Unmarshal on response bytes , %w", err)
			return
		}
		err = fmt.Errorf("[authorizeTemplateRequest] error occured in sending to orchestration service , %s", resObj.Message)
		return
	}
	var resobj responseobjects.AuthorizeResObj
	status = "client Side unmarshal response body"
	err = json.Unmarshal(bodyBytes, &resobj)
	if err != nil {
		err = fmt.Errorf("[authorizeTemplateRequest] error while trying to json unmarshal the bytes read from request body: %w", err)
		return
	}
	status = resobj.Status
	err = nil
	return

}
