package main

import (
	"runtime"

	"github.com/warnakulasuriya-fds-e23/fingerprint-go-client/uploader"
	"github.com/warnakulasuriya-fds-e23/go-sourceafis-fork/config"
)

func main() {
	config.LoadDefaultConfig()
	config.Config.Workers = runtime.NumCPU()
	// tests.Test8()
	uploader.Uploader()
}
