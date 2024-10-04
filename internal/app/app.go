package app

import (
	"fmt"
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
	a.dependencyProvider.Logger().Info(fmt.Sprintf("starting http server on -> %s", cfg.Addr()))

	return http.ListenAndServe(cfg.Addr(), a.initRoutes())
}
