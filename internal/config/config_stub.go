package config

type configDummy struct {
}

// NewConfigDummy конструктор stub конифга, только для тестов
func NewConfigDummy() Config {
	return &configDummy{}
}

// BaseReturnedURL стаб для полного урла
func (c *configDummy) BaseReturnedURL() string {
	return "localhost"
}

// FileStoragePath стаб для местоположения файла
func (c *configDummy) FileStoragePath() string {
	return "./url"
}

// DatabaseDSN стаб для DSN
func (c *configDummy) DatabaseDSN() string {
	return "postgres://postgres:postgres@localhost/postgres?sslmode=disable"
}

// EnableHTTPS стаб для https
func (c *configDummy) EnableHTTPS() bool {
	return false
}

// Addr стаб для адреса запускаемого сервера
func (c *configDummy) Addr() string {
	return ":8080"
}

// ConfigFile имя файла с конфигом
func (c *configDummy) ConfigFile() string {
	return "config"
}

// TrustedSubnet стаба для допускаемой подсети
func (c *configDummy) TrustedSubnet() string {
	return ""
}

// GRPCPort стаба для порта GRPC
func (c *configDummy) GRPCPort() string {
	return ""
}
