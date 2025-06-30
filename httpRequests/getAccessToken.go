package httprequests

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
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

func isTokenExpired(token string) {

}

func (client *Httpclientimpl) getAccessToken() string {
	client.mutex.Lock()
	defer client.mutex.Unlock()
	if client.accessToken == "" || client.expiryTime.Equal(time.Now()) || client.expiryTime.Before(time.Now().Add(5*time.Second)) {

		tokenEndpoint := os.Getenv("TOKEN_ENDPOINT")
		data := url.Values{}
		data.Set("grant_type", "client_credentials")
		requestBody := bytes.NewBufferString(data.Encode())
		req, err := http.NewRequest("POST", tokenEndpoint, requestBody)
		if err != nil {
			log.Fatal(err.Error())
		}
		consumerKey := os.Getenv("CONSUMMER_KEY")
		consumerSecret := os.Getenv("CONSUMER_SECRET")
		authHeadervalue := base64.StdEncoding.EncodeToString([]byte(consumerKey + ":" + consumerSecret))
		req.Header.Add("Authorization", "Basic "+authHeadervalue)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		internalclient := &http.Client{}
		res, err := internalclient.Do(req)
		if err != nil {
			log.Fatal(err.Error())
		}
		bodybytes, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err.Error())
		}

		var resObj tokenResponseObject
		err = json.Unmarshal(bodybytes, &resObj)
		if err != nil {
			log.Fatal(err.Error())
		}

		client.expiryTime = time.Now().Add(time.Duration(resObj.ExpiresIn) * time.Second)

		return resObj.AccessToken
	}

	return client.accessToken
}
