package app

import (
	"github.com/Makovey/shortener/internal/api"
	"github.com/Makovey/shortener/internal/api/shortener_api"
	"github.com/Makovey/shortener/internal/config"
	"github.com/Makovey/shortener/internal/logger"
	"github.com/Makovey/shortener/internal/logger/stdout"
	"github.com/Makovey/shortener/internal/repository"
	"github.com/Makovey/shortener/internal/repository/in_memory"
	"github.com/Makovey/shortener/internal/service"
	"github.com/Makovey/shortener/internal/service/shortener"
)

type dependencyProvider struct {
	shortHandler api.HttpHandler
	config       config.HttpConfig
	logger       logger.Logger

	shortRepo repository.ShortenerRepository
	shorSrv   service.ShortenerService
}

func newDependencyProvider() *dependencyProvider {
	return &dependencyProvider{}
}

func (p *dependencyProvider) HttpHandler() api.HttpHandler {
	if p.shortHandler == nil {
		p.shortHandler = shortener_api.NewShortenerHandler(p.ShortenerService(), p.Logger(), p.Config())
	}

	return p.shortHandler
}

func (p *dependencyProvider) Logger() logger.Logger {
	if p.logger == nil {
		p.logger = stdout.NewLoggerStdout()
	}

	return p.logger
}

func (p *dependencyProvider) ShortenerRepository() repository.ShortenerRepository {
	if p.shortRepo == nil {
		p.shortRepo = in_memory.NewRepositoryInMemory()
	}

	return p.shortRepo
}

func (p *dependencyProvider) ShortenerService() service.ShortenerService {
	if p.shorSrv == nil {
		p.shorSrv = shortener.NewShortenerService(p.ShortenerRepository())
	}

	return p.shorSrv
}

func (p *dependencyProvider) Config() config.HttpConfig {
	if p.config == nil {
		p.config = config.NewHttpConfig()
	}

	return p.config
}
