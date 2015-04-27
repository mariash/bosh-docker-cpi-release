package config

import (
	"encoding/json"

	boshsys "github.com/cloudfoundry/bosh-agent/system"
	bwcvm "github.com/cppforlife/bosh-warden-cpi/vm"
)

type Config struct {
	SocketPath string
	Mbus       string
	Blobstore  bwcvm.BlobstoreOptions
}

func LoadConfig(configPath string, fs boshsys.FileSystem) (Config, error) {
	var config Config

	contents, err := fs.ReadFile(configPath)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(contents, &config)

	return config, err
}
