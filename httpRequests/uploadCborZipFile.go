package httprequests

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

func (client *Httpclientimpl) uploadCborZipFile(zipFilePath string) error {
	var accessToken string
	if client.accessToken == "" || client.expiryTime.Equal(time.Now()) || client.expiryTime.Before(time.Now().Add(5*time.Second)) {
		accessToken = client.getAccessToken()
	} else {
		accessToken = client.accessToken
	}
	file, err := os.Open(zipFilePath)
	if err != nil {
		return fmt.Errorf("[uploadCborZipFile]failed to open file %s: %w", zipFilePath, err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", filepath.Base(zipFilePath))
	if err != nil {
		return fmt.Errorf("[uploadCborZipFile]failed to create form file: %w", err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return fmt.Errorf("[uploadCborZipFile]failed to copy file content to form: %w", err)
	}

	writer.Close()
	urlString, err := url.JoinPath(client.orchestrationServerAdress, UploadCborZipFileEndpoint)
	if err != nil {
		return fmt.Errorf("[uploadCborZipFile] error while trying to combine url %s with %s : %w", client.orchestrationServerAdress, MatchTemplatesEndpoint, err)
	}
	req, err := http.NewRequest("POST", urlString, body)
	if err != nil {
		return fmt.Errorf("[uploadCborZipFile]failed to create HTTP request: %w", err)
	}
	client.SetOrAddHeaderValueAccordingToKey("Content-Type", "multipart/form-data")
	client.SetOrAddHeaderValueAccordingToKey("Authorization", "Bearer "+accessToken)
	for _, headerKeyValuePair := range client.headerKeyValueArray {
		req.Header.Add(headerKeyValuePair.key, headerKeyValuePair.value)
	}

	internalClient := &http.Client{}
	resp, err := internalClient.Do(req)
	if err != nil {
		return fmt.Errorf("[uploadCborZipFile][matchTemplateRequest] error while sending match template http request or recieving response : %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("[uploadCborZipFile][matchTemplateRequest] error while reading bytes of response body : %w", err)
	}
	responseBodyString := string(bodyBytes)
	log.Println("recieved response: ", responseBodyString)
	return nil
}
