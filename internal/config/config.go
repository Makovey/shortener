package config

// Config methods:
// Addr - returned address of launching server
// BaseReturnedURL - returned base url in response when url is shorted
// FileStoragePath - returned path for file with urls
type Config interface {
	Addr() string
	BaseReturnedURL() string
	FileStoragePath() string
}

const (
	defaultAddr    = "localhost:8080"
	defaultBaseURL = "http://localhost:8080"
	defaultPathURL = "./urls.txt"
)

type config struct {
	addr            string
	baseReturnedURL string
	fileStoragePath string
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

func NewConfig() Config {
	envCfg := newEnvConfig()
	flags := newFlagsValue()

	return &config{
		addr:            addrValue(envCfg, flags),
		baseReturnedURL: baseURLValue(envCfg, flags),
		fileStoragePath: filePathValue(envCfg, flags),
	}
}

func addrValue(envCfg envConfig, flags flagsValue) string {
	var addr string
	if envAddr := envCfg.Addr; envAddr != "" {
		addr = envAddr
	} else if flagAddr := flags.addr; flagAddr != "" {
		addr = flags.addr
	} else {
		addr = defaultAddr
	}

	return addr
}

func baseURLValue(envCfg envConfig, flags flagsValue) string {
	var baseReturnedURL string
	if envBase := envCfg.BaseReturnedURL; envBase != "" {
		baseReturnedURL = envBase
	} else if flagBaseURL := flags.baseReturnedURL; flagBaseURL != "" {
		baseReturnedURL = flags.baseReturnedURL
	} else {
		baseReturnedURL = defaultBaseURL
	}

	return baseReturnedURL
}

func filePathValue(envCfg envConfig, flags flagsValue) string {
	var urlPath string
	if envStoragePath := envCfg.FileStoragePath; envStoragePath != "" {
		urlPath = envStoragePath
	} else if flagPath := flags.fileStoragePath; flagPath != "" {
		urlPath = flags.fileStoragePath
	} else {
		urlPath = defaultPathURL
	}

	return urlPath
}
