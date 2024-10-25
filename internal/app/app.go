package app

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"

	"github.com/Makovey/shortener/internal/middleware"
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
	r.Use(middleware.NewCompressor().Compress)
	r.Use(chiMiddleware.Recoverer)

	r.Post("/", a.dependencyProvider.HTTPHandler().PostNewURLHandler)
	r.Post("/api/shorten", a.dependencyProvider.HTTPHandler().PostShortenURLHandler)
	r.Get("/{id}", a.dependencyProvider.HTTPHandler().GetURLHandler)

	return r
}

func (a *App) runHTTPServer() {
	cfg := a.dependencyProvider.Config()
	a.dependencyProvider.Logger().Info(fmt.Sprintf("starting http server on -> %s", cfg.Addr()))

	if err := http.ListenAndServe(cfg.Addr(), a.initRouter()); err != nil {
		a.dependencyProvider.Logger().Info(fmt.Sprintf("http server stopped: %s", err))
	}
}
