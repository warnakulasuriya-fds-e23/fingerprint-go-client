package main

import (
	"runtime"

	"github.com/warnakulasuriya-fds-e23/fingerprint-go-client/tests"
	"github.com/warnakulasuriya-fds-e23/go-sourceafis-fork/config"
)

func main() {
	config.LoadDefaultConfig()
	config.Config.Workers = runtime.NumCPU()
	tests.Test1()
}
