package config

const defaultPort = ":8080"

type HTTPConfig interface {
	Port() string
}

type httpConfig struct {
	port string
}

func (cfg *httpConfig) Port() string {
	return cfg.port
}

func NewHTTPConfig() HTTPConfig {
	return &httpConfig{
		port: defaultPort,
	}
}
