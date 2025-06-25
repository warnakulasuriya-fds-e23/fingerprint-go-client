package main

import (
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-client/configtomlreader"
	httprequests "github.com/warnakulasuriya-fds-e23/fingerprint-go-client/httpRequests"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-client/uploader"
)

func main() {
	client := httprequests.NewHttpClientImpl()
	config := configtomlreader.ConfigTomlReader()
	client.SetOrAddHeaderValueAccordingToKey("Test-Key", config.TestKey)

	// tests.Test8()
	uploader.Uploader(client)
}
