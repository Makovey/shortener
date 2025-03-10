package main

import (
	"database/sql/driver"
	"fmt"
	"io"
	syslog "log"

	"github.com/Makovey/shortener/internal/app"
	"github.com/Makovey/shortener/internal/closer"
	"github.com/Makovey/shortener/internal/config"
	"github.com/Makovey/shortener/internal/logger"
	"github.com/Makovey/shortener/internal/logger/stdout"
	"github.com/Makovey/shortener/internal/repository/disc"
	"github.com/Makovey/shortener/internal/repository/inmemory"
	"github.com/Makovey/shortener/internal/repository/postgres"
	"github.com/Makovey/shortener/internal/service/shortener"
	"github.com/Makovey/shortener/internal/transport/grpc/service_info"
	shortener_server "github.com/Makovey/shortener/internal/transport/grpc/shortener"
	transport "github.com/Makovey/shortener/internal/transport/http"
)

var (
	buildVersion = "N/A" // версия приложения
	buildDate    = "N/A" // дата сборки
	buildCommit  = "N/A" // коммит сборки
)

func main() {
	log := stdout.NewLoggerStdout(stdout.EnvLocal)
	cfg := config.NewConfig(log)
	closers := closer.NewCloser()

	repo := assembleRepo(cfg, log, closers)
	pinger := assemblePinger(repo, cfg, log, closers)
	checker := shortener.NewChecker(pinger)
	service := shortener.NewShortenerService(repo, cfg, log)

	handler := transport.NewHTTPHandler(
		service,
		log,
		checker,
	)

	infoServer := service_info.NewInfoServer(log, checker, service)
	shortenerServer := shortener_server.NewServer(log, service)

	appl := app.NewApp(
		log,
		cfg,
		handler,
		infoServer,
		shortenerServer,
	)

	log.Debug(fmt.Sprintf("Build version: %s", buildVersion))
	log.Debug(fmt.Sprintf("Build date: %s", buildDate))
	log.Debug(fmt.Sprintf("Build commit: %s", buildCommit))

	appl.Run()

	defer closers.CloseAll()
}

func assembleRepo(
	cfg config.Config,
	log logger.Logger,
	closer *closer.Closer,
) shortener.Repository {
	var repo shortener.Repository
	switch {
	case cfg.DatabaseDSN() != "":
		postgre, err := postgres.NewPostgresRepository(cfg, log)
		if err != nil {
			syslog.Fatalf("unable to create postgres repository: %s", err)
		}
		repo = postgre
	case cfg.FileStoragePath() != "":
		file, err := disc.NewFileStorage(cfg.FileStoragePath(), log)
		if err != nil {
			syslog.Fatalf("unable to create file storage repository: %s", err)
		}
		repo = file
	default:
		repo = inmemory.NewRepositoryInMemory()
	}

	if c, ok := repo.(io.Closer); ok {
		closer.Add(c)
	}

	return repo
}

func assemblePinger(
	repo shortener.Repository,
	cfg config.Config,
	log logger.Logger,
	closers *closer.Closer,
) driver.Pinger {
	var pinger driver.Pinger

	if ping, ok := repo.(driver.Pinger); ok {
		pinger = ping
	} else {
		postgresPing, err := postgres.NewPingerRepo(cfg)
		if err != nil {
			syslog.Fatalf("unable to create postgres pinger: %s", err)
		}
		pinger = postgresPing
	}

	if c, ok := pinger.(io.Closer); ok {
		closers.Add(c)
	}

	return pinger
}
