package config

// HTTPConfig methods:
// Addr - returned address of launching server
// BaseReturnedURL - returned base url in response when url is shorted
type HTTPConfig interface {
	Addr() string
	BaseReturnedURL() string
}

type httpConfig struct {
	addr            string
	baseReturnedURL string
}

func (cfg *httpConfig) Addr() string {
	return cfg.addr
}

func (cfg *httpConfig) BaseReturnedURL() string {
	return cfg.baseReturnedURL
}

func NewHTTPConfig() HTTPConfig {
	flags := newFlagsValue()

	return &httpConfig{
		addr:            flags.addr,
		baseReturnedURL: flags.baseReturnedURL,
	}
}
