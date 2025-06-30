package httprequests

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

type tokenResponseObject struct {
	AccessToken      string `json:"access_token"`
	TokenType        string `json:"token_type"`
	ExpiresIn        int    `json:"expires_in"`
	Scope            string `json:"scope,omitempty"`
	Error            string `json:"error,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

func (client *Httpclientimpl) getAccessToken() (token string, err error) {
	client.mutex.Lock()
	defer client.mutex.Unlock()
	if client.accessToken == "" || client.expiryTime.Equal(time.Now()) || client.expiryTime.Before(time.Now().Add(5*time.Second)) {

		tokenEndpoint := os.Getenv("TOKEN_ENDPOINT")
		data := url.Values{}
		data.Set("grant_type", "client_credentials")
		requestBody := bytes.NewBufferString(data.Encode())
		req, errNewReq := http.NewRequest("POST", tokenEndpoint, requestBody)
		if errNewReq != nil {
			token = ""
			err = fmt.Errorf("error while creating a post request for the tokenEndpoint : %w", errNewReq)
			return
		}
		consumerKey := os.Getenv("CONSUMMER_KEY")
		consumerSecret := os.Getenv("CONSUMER_SECRET")
		authHeadervalue := base64.StdEncoding.EncodeToString([]byte(consumerKey + ":" + consumerSecret))
		req.Header.Add("Authorization", "Basic "+authHeadervalue)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		internalclient := &http.Client{}
		res, errReqSend := internalclient.Do(req)
		if errReqSend != nil {
			token = ""
			err = fmt.Errorf("error while sending or recieving post request : %w", errReqSend)
			return
		}
		bodybytes, errReadAll := io.ReadAll(res.Body)
		if errReadAll != nil {
			token = ""
			err = fmt.Errorf("error while reading bytes of response body : %w", errReadAll)
			return
		}

		var resObj tokenResponseObject
		errUnMarshal := json.Unmarshal(bodybytes, &resObj)
		if errUnMarshal != nil {
			token = ""
			err = fmt.Errorf("error while running json unmarshal for the read bytes of the response body : %w", errUnMarshal)
			return
		}

		client.expiryTime = time.Now().Add(time.Duration(resObj.ExpiresIn) * time.Second)
		client.accessToken = resObj.AccessToken

	}

	token = client.accessToken
	err = nil
	return

}
