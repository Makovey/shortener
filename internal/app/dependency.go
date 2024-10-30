package app

import (
	"github.com/Makovey/shortener/internal/api"
	"github.com/Makovey/shortener/internal/api/shortenerapi"
	"github.com/Makovey/shortener/internal/config"
	"github.com/Makovey/shortener/internal/logger"
	"github.com/Makovey/shortener/internal/logger/stdout"
	"github.com/Makovey/shortener/internal/repository/file"
	"github.com/Makovey/shortener/internal/service"
	"github.com/Makovey/shortener/internal/service/shortener"
)

type dependencyProvider struct {
	shortHandler api.HTTPHandler
	config       config.Config
	logger       logger.Logger

	shortRepo service.Shortener
	shorSrv   api.Shortener
}

func newDependencyProvider() *dependencyProvider {
	return &dependencyProvider{}
}

func (p *dependencyProvider) HTTPHandler() api.HTTPHandler {
	if p.shortHandler == nil {
		p.shortHandler = shortenerapi.NewShortenerHandler(p.ShortenerService(), p.Logger(), p.Config())
	}

	return p.shortHandler
}

func (p *dependencyProvider) Logger() logger.Logger {
	if p.logger == nil {
		p.logger = stdout.NewLoggerStdout("local") // TODO: Add env deployment config
	}

	return p.logger
}

func (p *dependencyProvider) ShortenerRepository() service.Shortener {
	if p.shortRepo == nil {
		p.shortRepo = file.NewFileStorage(p.config.FileStoragePath(), p.Logger())
	}

	return p.shortRepo
}

func (p *dependencyProvider) ShortenerService() api.Shortener {
	if p.shorSrv == nil {
		p.shorSrv = shortener.NewShortenerService(p.ShortenerRepository())
	}

	return p.shorSrv
}

func (p *dependencyProvider) Config() config.Config {
	if p.config == nil {
		p.config = config.NewConfig()
	}

	return p.config
}
