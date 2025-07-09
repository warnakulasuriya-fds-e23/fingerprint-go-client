package configtomlreader

import (
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-client/internal/configuration"
)

func ConfigTomlReader() configuration.ConfigurationToml {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}

	tomlpath := filepath.Join(workingDir, "config.toml")
	var config configuration.ConfigurationToml
	toml.DecodeFile(tomlpath, &config)
	return config
}
