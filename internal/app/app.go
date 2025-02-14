// Package app, ответственный за запуск веб-сервера, его выключение и роутинг хендлеров.
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

	"golang.org/x/crypto/acme/autocert"

	"github.com/Makovey/shortener/internal/config"
	"github.com/Makovey/shortener/internal/logger"
	"github.com/Makovey/shortener/internal/transport"
)

// App содержит в себе зависимости, необходимоые для запуска веб-сервера и его корректной работы.
type App struct {
	log     logger.Logger         // для логирования дополнительной информации
	cfg     config.Config         // конфиг, в котором лежит адрес, на котором будет запущен сервер
	handler transport.HTTPHandler // хэндлеры HTTP-запросов
	wg      sync.WaitGroup        // для синхронизации горутин
}

// NewApp конструктор App
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

// Run запуск всех процессов, которыми владеет App.
func (a *App) Run() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGTERM)
	defer stop()

	a.runHTTPServer(ctx)

	a.wg.Wait()
}

// runHTTPServer запускает HTTP сервер.
func (a *App) runHTTPServer(ctx context.Context) {
	fn := "app.runHTTPServer"

	a.wg.Add(1)
	defer a.wg.Done()

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS13,
	}

	if a.cfg.EnableHTTPS() {
		manager := &autocert.Manager{
			Cache:      autocert.DirCache("cache-dir"),
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist("localhost"),
		}

		tlsConfig = manager.TLSConfig()
	}

	srv := &http.Server{
		Addr:      a.cfg.Addr(),
		Handler:   a.initRouter(),
		TLSConfig: tlsConfig,
	}

	a.log.Info(fmt.Sprintf("[%s]: starting http server on: %s", fn, a.cfg.Addr()))

	go func() {
		if a.cfg.EnableHTTPS() {
			if err := srv.ListenAndServeTLS("", ""); err != nil {
				a.log.Info(fmt.Sprintf("[%s] http server stopped: %s", fn, err))
			}
		} else {
			if err := srv.ListenAndServe(); err != nil {
				a.log.Info(fmt.Sprintf("[%s] http server stopped: %s", fn, err))
			}
		}
	}()

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		a.log.Info(fmt.Sprintf("[%s]: http server shutdown timeout: %s", err, fn))
	}
}
