package config

const defaultPort = ":8080"

type HttpConfig interface {
	Port() string
}

type httpConfig struct {
	port string
}

func (cfg *httpConfig) Port() string {
	return cfg.port
}

func NewHttpConfig() HttpConfig {
	return &httpConfig{
		port: defaultPort,
	}
}
