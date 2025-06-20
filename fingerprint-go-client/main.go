package main

import (
	"runtime"

	"github.com/miqdadyyy/go-sourceafis/config"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-client/tests"
)

func main() {
	config.LoadDefaultConfig()
	config.Config.Workers = runtime.NumCPU()
	tests.Test1()
}
