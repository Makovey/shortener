package config

import (
	"log"

	"github.com/caarlos0/env/v6"
)

type envConfig struct {
	Addr            string `env:"SERVER_ADDRESS"`
	BaseReturnedURL string `env:"BASE_URL"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
}

func newEnvConfig() envConfig {
	var cfg envConfig
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}
