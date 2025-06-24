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

func (client *httpclientimpl) enrollTemplateRequest(newEntry *templates.SearchTemplate, id string) (message string, err error) {
	message = ""
	newEntryData, err := client.sdk.GetAsByteArray(newEntry)
	if err != nil {
		err = fmt.Errorf("error occured while trying to convert newEntry template to Byte array, %w", err)
		return
	}
	reqObj := requestobjects.EnrollTemplateReqObj{Data: *newEntryData, Id: id}
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
	resp, err := http.Post(urlString, "application/json", bytes.NewBuffer(jsonobj))
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
