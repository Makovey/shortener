package config

import (
	"log"

	"github.com/caarlos0/env/v6"
)

type envConfig struct {
	Addr            string `env:"SERVER_ADDRESS"`
	BaseReturnedURL string `env:"BASE_URL"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
	DatabaseDSN     string `env:"DATABASE_DSN"`
	EnableHTTPS     bool   `env:"ENABLE_HTTPS"`
	ConfigFilePath  string `env:"CONFIG"`
	TrustedSubnet   string `env:"TRUSTED_SUBNET"`
	GRPCPort        string `env:"GRPC_PORT"`
}

func newEnvConfig() envConfig {
	var cfg envConfig
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}
