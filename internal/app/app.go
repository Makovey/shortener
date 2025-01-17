package app

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/Makovey/shortener/internal/config"
	"github.com/Makovey/shortener/internal/logger"
	"github.com/Makovey/shortener/internal/transport"
)

type App struct {
	log     logger.Logger
	cfg     config.Config
	handler transport.HTTPHandler
	wg      sync.WaitGroup
}

func NewApp(
	log logger.Logger,
	cfg config.Config,
	handler transport.HTTPHandler,
) *App {
	return &App{
		log:     log,
		cfg:     cfg,
		handler: handler,
		wg:      sync.WaitGroup{},
	}
}

func (a *App) Run() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGTERM)
	defer stop()

	a.runHTTPServer(ctx)

	a.wg.Wait()
}

func (a *App) runHTTPServer(ctx context.Context) {
	a.wg.Add(1)
	defer a.wg.Done()

	fn := "app.runHTTPServer"

	a.log.Info(fmt.Sprintf("[%s]: starting http server on: %s", fn, a.cfg.Addr()))

	srv := &http.Server{
		Addr:    a.cfg.Addr(),
		Handler: a.initRouter(),
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS13,
		},
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			a.log.Info(fmt.Sprintf("[%s] http server stopped: %s", fn, err))
		}
	}()

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		a.log.Info(fmt.Sprintf("[%s]: http server shutdown timeout: %s", err, fn))
	}
}
