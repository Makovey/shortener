package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type App struct {
	dependencyProvider *dependencyProvider
}

func NewApp() *App {
	return &App{dependencyProvider: newDependencyProvider()}
}

func (a *App) Run() error {
	return a.runHTTPServer()
}

func (a *App) initRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/", a.dependencyProvider.HTTPHandler().PostNewURLHandler)
	r.Get("/{id}", a.dependencyProvider.HTTPHandler().GetURLHandler)

	return r
}

func (a *App) runHTTPServer() error {
	cfg := a.dependencyProvider.Config()

	return http.ListenAndServe(cfg.Port(), a.initRoutes())
}
