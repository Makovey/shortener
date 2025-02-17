package config

import (
	"encoding/json"
	"os"
)

type jsonConfig struct {
	Addr            string `json:"server_address"`
	BaseURL         string `json:"base_url"`
	FileStoragePath string `json:"file_storage_path"`
	DatabaseDSN     string `json:"database_dsn"`
	EnableHTTPS     bool   `json:"enable_https"`
}

func newJSONFileConfig(filePath string) jsonConfig {
	var config jsonConfig

	file, err := os.Open(filePath + ".json")
	if err != nil {
		return jsonConfig{}
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return jsonConfig{}
	}

	return config
}
