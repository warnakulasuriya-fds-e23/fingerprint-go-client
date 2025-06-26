package tests

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type tokenResponseObject struct {
	AccessToken      string `json:"access_token"`
	TokenType        string `json:"token_type"`
	ExpiresIn        int    `json:"expires_in"`
	Scope            string `json:"scope,omitempty"`
	Error            string `json:"error,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

func Test9() {
	tokenEndpoint := "https://3ee2e1ea-a3ec-4f5b-bff1-2fa3a86782c0-prod.e1-us-east-azure.choreosts.dev/oauth2/token"
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
	fmt.Println(authHeadervalue)
	req.Header.Add("Authorization", "Basic "+authHeadervalue)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}
	bodybytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	responseBodyString := string(bodybytes)
	fmt.Println(responseBodyString)

	var resObj tokenResponseObject
	err = json.Unmarshal(bodybytes, &resObj)
	if err != nil {
		log.Fatal(err.Error())
	}
	println(resObj.AccessToken)
}
