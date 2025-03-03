// Package config отвечат за конфигурацию приложения посредством env и flag переменных.
package config

import (
	"strconv"
	"strings"

	"github.com/Makovey/shortener/internal/logger"
)

// Config аггрегирует конфигурацию через флаги и переменные окружения.
// Приоритет: Env -> Flag -> ConfigFile -> Default
type Config interface {
	Addr() string            // returned address of launching server
	BaseReturnedURL() string // returned base url in response when url is shorted
	FileStoragePath() string // returned path for disc with urls
	DatabaseDSN() string     // data source string for sql.DB
	EnableHTTPS() bool       // enabled https
	ConfigFile() string      // name of config file
	TrustedSubnet() string   // returned trusted subnets for handler
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
	enableHTTPS     bool
	configFile      string
	trustedSubnet   string
}

// NewConfig конструктор Config
func NewConfig(
	log logger.Logger, // необходимо для логирования конфигурации на старте приложения
) Config {
	envCfg := newEnvConfig()
	flags := newFlagsValue()
	file := parseJSONConfig(envCfg.ConfigFilePath, flags.configFilePath)

	cfg := &config{
		addr:            resolveValue(envCfg.Addr, flags.addr, file.Addr, defaultAddr),
		baseReturnedURL: resolveValue(envCfg.BaseReturnedURL, flags.baseReturnedURL, file.BaseURL, defaultBaseURL),
		fileStoragePath: resolveValue(envCfg.FileStoragePath, flags.fileStoragePath, file.FileStoragePath, ""),
		databaseDSN:     resolveDatabaseURI(envCfg.DatabaseDSN, flags.databaseDSN, file.DatabaseDSN),
		enableHTTPS:     resolveBoolValue(envCfg.EnableHTTPS, flags.enableHTTPS, file.EnableHTTPS),
	}

	log.Debug("Addr: " + cfg.addr)
	log.Debug("BaseReturnedURL: " + cfg.baseReturnedURL)
	log.Debug("FileStoragePath: " + cfg.fileStoragePath)
	log.Debug("DatabaseDSN: " + cfg.databaseDSN)
	log.Debug("EnableHTTPS: " + strconv.FormatBool(cfg.enableHTTPS))

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

// EnableHTTPS - поднятие сервера под HTTPS
func (cfg config) EnableHTTPS() bool {
	return cfg.enableHTTPS
}

// ConfigFile - путь до файла конфигурации
func (cfg config) ConfigFile() string {
	return cfg.fileStoragePath
}

// TrustedSubnet - подсеть с которой разрешено получать стату
func (cfg config) TrustedSubnet() string {
	return cfg.trustedSubnet
}

func resolveValue(envValue, flagValue, fileValue, defaultValue string) string {
	if envValue != "" {
		return envValue
	}

	if flagValue != "" {
		return flagValue
	}

	if fileValue != "" {
		return fileValue
	}

	return defaultValue
}

func resolveDatabaseURI(envValue, flagValue, fileValue string) string {
	dsn := resolveValue(envValue, flagValue, fileValue, "")
	if dsn != "" && !strings.Contains(dsn, "?sslmode=disable") {
		dsn += "?sslmode=disable"
	}
	return dsn
}

func resolveBoolValue(envValue, flagValue, fileValue bool) bool {
	if envValue || flagValue || fileValue {
		return true
	}

	return false
}

func parseJSONConfig(envPath, flagPath string) jsonConfig {
	var path string

	if envPath != "" {
		path = envPath
	}

	if flagPath != "" {
		path = flagPath
	}

	return newJSONFileConfig(path)
}
