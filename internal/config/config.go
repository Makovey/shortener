package config

import (
	"strings"

	"github.com/Makovey/shortener/internal/logger"
)

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

func NewConfig(
	log logger.Logger,
) Config {
	envCfg := newEnvConfig()
	flags := newFlagsValue()

	cfg := &config{
		addr:            resolveValue(envCfg.Addr, flags.addr, defaultAddr),
		baseReturnedURL: resolveValue(envCfg.BaseReturnedURL, flags.baseReturnedURL, defaultBaseURL),
		fileStoragePath: resolveValue(envCfg.FileStoragePath, flags.fileStoragePath, ""),
		databaseDSN:     resolveDatabaseURI(envCfg.DatabaseDSN, flags.databaseDSN),
	}

	log.Debug("Addr: " + cfg.addr)
	log.Debug("BaseReturnedURL: " + cfg.addr)
	log.Debug("FileStoragePath: " + cfg.addr)
	log.Debug("DatabaseDSN: " + cfg.addr)

	return cfg
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

func resolveValue(envValue, flagValue, defaultValue string) string {
	if envValue != "" {
		return envValue
	}

	if flagValue != "" {
		return flagValue
	}

	return defaultValue
}

func resolveDatabaseURI(envValue, flagValue string) string {
	dsn := resolveValue(envValue, flagValue, "")
	if dsn != "" && !strings.Contains(dsn, "?sslmode=disable") {
		dsn += "?sslmode=disable"
	}
	return dsn
}
