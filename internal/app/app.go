package app

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"

	"github.com/Makovey/shortener/internal/middleware"
	"github.com/Makovey/shortener/internal/middleware/utils"
)

type App struct {
	dependencyProvider *dependencyProvider
}

func NewApp() *App {
	return &App{dependencyProvider: newDependencyProvider()}
}

func (a *App) Run() {
	a.runHTTPServer()
}

func (a *App) initRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.NewLogger(a.dependencyProvider.Logger()).Logger)

	jwtMiddleware := middleware.NewAuthHandler(a.dependencyProvider.Logger(), utils.NewJWTUtils(a.dependencyProvider.Logger()))
	r.Use(jwtMiddleware.AuthHandler)
	r.Use(middleware.NewCompressor().Compress)
	r.Use(chiMiddleware.Recoverer)

	r.Post("/", a.dependencyProvider.HTTPHandler().PostNewURL)
	r.Get("/{id}", a.dependencyProvider.HTTPHandler().GetURL)
	r.Get("/ping", a.dependencyProvider.HTTPHandler().GetPing)

	r.Route("/api", func(r chi.Router) {
		r.Post("/shorten", a.dependencyProvider.HTTPHandler().PostShortenURL)
		r.Post("/shorten/batch", a.dependencyProvider.HTTPHandler().PostBatch)
		r.Get("/user/urls", a.dependencyProvider.HTTPHandler().GetAllURLS)
		r.Delete("/user/urls", a.dependencyProvider.HTTPHandler().DeleteURLS)
	})

	return r
}

func (a *App) runHTTPServer() {
	cfg := a.dependencyProvider.Config()
	a.dependencyProvider.Logger().Info(fmt.Sprintf("starting http server on -> %s", cfg.Addr()))

	srv := &http.Server{
		Addr:    cfg.Addr(),
		Handler: a.initRouter(),
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS13,
		},
	}

	if err := srv.ListenAndServe(); err != nil {
		a.dependencyProvider.Logger().Info(fmt.Sprintf("http server stopped: %s", err))
	}

	defer a.dependencyProvider.Closer.CloseAll()
}
