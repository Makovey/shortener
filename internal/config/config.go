// Package config отвечат за конфигурацию приложения посредством env и flag переменных.
package config

import (
	"strings"

	"github.com/Makovey/shortener/internal/logger"
)

// Config аггрегирует конфигурацию через флаги и переменные окружения.
// Приоритет: Env -> Flag -> Default
type Config interface {
	Addr() string            // returned address of launching server
	BaseReturnedURL() string // returned base url in response when url is shorted
	FileStoragePath() string // returned path for disc with urls
	DatabaseDSN() string     // data source string for sql.DB
}

// Настройки по-умолчанию
const (
	defaultAddr    = "localhost:8080"        // запуск сервера
	defaultBaseURL = "http://localhost:8080" // возвращаемый base url для ответа некоторых ручек
)

type config struct {
	addr            string
	baseReturnedURL string
	fileStoragePath string
	databaseDSN     string
}

// NewConfig конструктор Config
func NewConfig(
	log logger.Logger, // необходимо для логирования конфигурации на старте приложения
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
	log.Debug("BaseReturnedURL: " + cfg.baseReturnedURL)
	log.Debug("FileStoragePath: " + cfg.fileStoragePath)
	log.Debug("DatabaseDSN: " + cfg.databaseDSN)

	return cfg
}

// Addr - сконфигурированный адрес для запуска сервера
func (cfg config) Addr() string {
	return cfg.addr
}

// BaseReturnedURL - сконфигурированный адрес в ответе для некоторых ручек
func (cfg config) BaseReturnedURL() string {
	return cfg.baseReturnedURL
}

// FileStoragePath - адрес файла на диске, как временное хранилище
func (cfg config) FileStoragePath() string {
	return cfg.fileStoragePath
}

// DatabaseDSN - адрес подключения к базе данных
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
