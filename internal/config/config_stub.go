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

// Addr стаб для адреса запускаемого сервера
func (c *configDummy) Addr() string {
	return ":8080"
}
