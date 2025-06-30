package httprequests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/warnakulasuriya-fds-e23/fingerprint-go-client/internal/requestobjects"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-client/internal/responseobjects"
	"github.com/warnakulasuriya-fds-e23/go-sourceafis-fork/templates"
)

func (client *Httpclientimpl) enrollTemplateRequest(newEntry *templates.SearchTemplate, id string) (message string, err error) {
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
	message = ""
	newEntryData, err := client.sdk.GetAsByteArray(newEntry)
	if err != nil {
		err = fmt.Errorf("error occured while trying to convert newEntry template to Byte array, %w", err)
		return
	}
	reqObj := requestobjects.EnrollTemplateReqObj{Data: *newEntryData, Id: id, ClientId: os.Getenv("DEVICE_ID")}
	jsonobj, err := json.Marshal(reqObj)
	if err != nil {
		err = fmt.Errorf("error occured while trying to convert reqObj to json, %w", err)
		return
	}
	urlString, err := url.JoinPath(client.orchestrationServerAdress, EnrollTemplateEndpoint)
	if err != nil {
		err = fmt.Errorf("error occured while trying to url string using url.JoinPath , %w", err)
		return
	}
	requestBody := bytes.NewBuffer(jsonobj)
	req, err := http.NewRequest("POST", urlString, requestBody)
	if err != nil {
		log.Fatal(err.Error())
	}

	client.SetOrAddHeaderValueAccordingToKey("Authorization", "Bearer "+accessToken)
	for _, headerKeyValuePair := range client.headerKeyValueArray {
		req.Header.Add(headerKeyValuePair.key, headerKeyValuePair.value)
	}

	internalClient := &http.Client{}
	resp, err := internalClient.Do(req)
	if err != nil {
		err = fmt.Errorf("error occured while using http.Post , %w", err)
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("error occured while reading response bytes using io.ReadAll , %w", err)
		return
	}
	var resObj responseobjects.EnrollTemplateResObj
	err = json.Unmarshal(bodyBytes, &resObj)
	if err != nil {
		err = fmt.Errorf("error occured while runnig json.Unmarshal on response bytes , %w", err)
		return
	}
	message = resObj.Message
	err = nil
	return
}
