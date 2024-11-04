package config

import "strings"

// Config methods:
// Addr - returned address of launching server
// BaseReturnedURL - returned base url in response when url is shorted
// FileStoragePath - returned path for disc with urls
// DatabaseDSN - data source string for sql.DB
type Config interface {
	Addr() string
	BaseReturnedURL() string
	FileStoragePath() string
	DatabaseDSN() string
}

const (
	defaultAddr    = "localhost:8080"
	defaultBaseURL = "http://localhost:8080"
)

type config struct {
	addr            string
	baseReturnedURL string
	fileStoragePath string
	databaseDSN     string
}

func (cfg config) Addr() string {
	return cfg.addr
}

func (cfg config) BaseReturnedURL() string {
	return cfg.baseReturnedURL
}

func (cfg config) FileStoragePath() string {
	return cfg.fileStoragePath
}

func (cfg config) DatabaseDSN() string {
	return cfg.databaseDSN
}

func NewConfig() Config {
	envCfg := newEnvConfig()
	flags := newFlagsValue()

	return &config{
		addr:            addrValue(envCfg, flags),
		baseReturnedURL: baseURLValue(envCfg, flags),
		fileStoragePath: filePathValue(envCfg, flags),
		databaseDSN:     databaseDSNValue(envCfg, flags),
	}
}

func addrValue(envCfg envConfig, flags flagsValue) string {
	addr := defaultAddr
	if envAddr := envCfg.Addr; envAddr != "" {
		addr = envAddr
	} else if flagAddr := flags.addr; flagAddr != "" {
		addr = flags.addr
	}

	return addr
}

func baseURLValue(envCfg envConfig, flags flagsValue) string {
	baseReturnedURL := defaultBaseURL
	if envBase := envCfg.BaseReturnedURL; envBase != "" {
		baseReturnedURL = envBase
	} else if flagBaseURL := flags.baseReturnedURL; flagBaseURL != "" {
		baseReturnedURL = flags.baseReturnedURL
	}

	return baseReturnedURL
}

func filePathValue(envCfg envConfig, flags flagsValue) string {
	var urlPath string
	if envStoragePath := envCfg.FileStoragePath; envStoragePath != "" {
		urlPath = envStoragePath
	} else if flagPath := flags.fileStoragePath; flagPath != "" {
		urlPath = flags.fileStoragePath
	}

	return urlPath
}

func databaseDSNValue(envCfg envConfig, flags flagsValue) string {
	var databaseDSN string
	if envDSN := envCfg.DatabaseDSN; envDSN != "" {
		databaseDSN = envDSN
	} else if flagDSN := flags.databaseDSN; flagDSN != "" {
		databaseDSN = flagDSN
	}

	if databaseDSN != "" && !strings.Contains(databaseDSN, "?sslmode=disable") {
		databaseDSN = databaseDSN + "?sslmode=disable"
	}

	return databaseDSN
}
