// Package app, ответственный за запуск веб-сервера, его выключение и роутинг хендлеров.
package app

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"golang.org/x/crypto/acme/autocert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	_ "google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/reflection"

	"github.com/Makovey/shortener/internal/config"
	protoInfo "github.com/Makovey/shortener/internal/generated/service_info"
	protoShortener "github.com/Makovey/shortener/internal/generated/shortener"
	"github.com/Makovey/shortener/internal/interceptor"
	"github.com/Makovey/shortener/internal/logger"
	"github.com/Makovey/shortener/internal/middleware/utils"
	"github.com/Makovey/shortener/internal/transport"
	"github.com/Makovey/shortener/internal/transport/grpc/service_info"
	"github.com/Makovey/shortener/internal/transport/grpc/shortener"
)

// App содержит в себе зависимости, необходимоые для запуска веб-сервера и его корректной работы.
type App struct {
	log             logger.Logger         // для логирования дополнительной информации
	cfg             config.Config         // конфиг, в котором лежит адрес, на котором будет запущен сервер
	handler         transport.HTTPHandler // хэндлеры HTTP-запросов
	infoServer      *service_info.InfoServer
	shortenerServer *shortener.Server
}

// NewApp конструктор App
func NewApp(
	log logger.Logger,
	cfg config.Config,
	handler transport.HTTPHandler,
	infoServer *service_info.InfoServer,
	shortenerServer *shortener.Server,
) *App {
	return &App{
		log:             log,
		cfg:             cfg,
		handler:         handler,
		infoServer:      infoServer,
		shortenerServer: shortenerServer,
	}
}

// Run запуск всех процессов, которыми владеет App.
func (a *App) Run() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup

	wg.Add(1)
	go a.runHTTPServer(ctx, &wg)

	wg.Add(1)
	go a.runGRPCServer(ctx, &wg)

	wg.Wait()
}

// runHTTPServer запускает HTTP сервер.
func (a *App) runHTTPServer(ctx context.Context, wg *sync.WaitGroup) {
	fn := "app.runHTTPServer"

	defer wg.Done()

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

func (a *App) runGRPCServer(ctx context.Context, wg *sync.WaitGroup) {
	fn := "app.runGRPCServer"

	defer wg.Done()

	listen, err := net.Listen("tcp", a.cfg.GRPCPort())
	if err != nil {
		a.log.Error(fmt.Sprintf("[%s]: failed to listen: %s", fn, err.Error()))
		return
	}

	s := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.ChainUnaryInterceptor(
			interceptor.Logger(a.log),
			interceptor.JWTAuth(a.log, utils.NewJWTUtils(a.log)),
			interceptor.CheckSubnet(a.cfg.TrustedSubnet(), []string{
				"Stats",
			}),
		),
	)

	reflection.Register(s)
	protoInfo.RegisterServiceInfoServer(s, a.infoServer)
	protoShortener.RegisterShortenerServer(s, a.shortenerServer)

	a.log.Info(fmt.Sprintf("[%s]: starting grpc server on: %s", fn, a.cfg.GRPCPort()))

	go func() {
		if err = s.Serve(listen); err != nil {
			a.log.Error(fmt.Sprintf("[%s]: can't serve grpc server: %s", fn, err.Error()))
		}
	}()

	<-ctx.Done()

	shutDownCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	isStoppedCh := make(chan struct{})
	go func() {
		s.GracefulStop()
		close(isStoppedCh)
	}()

	select {
	case <-isStoppedCh:
		a.log.Info(fmt.Sprintf("[%s]: grpc server stopped gracefully", fn))
	case <-shutDownCtx.Done():
		a.log.Error(fmt.Sprintf("[%s]: graceful shutdown timeout reached, forcing shutdown", fn))
		s.Stop()
	}
}
