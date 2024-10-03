package app

import "net/http"

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
	mux := http.NewServeMux()
	mux.HandleFunc("POST /", a.dependencyProvider.HTTPHandler().PostNewURLHandler)
	mux.HandleFunc("GET /{id}", a.dependencyProvider.HTTPHandler().GetURLHandler)

	return mux
}

func (a *App) runHTTPServer() error {
	cfg := a.dependencyProvider.Config()

	return http.ListenAndServe(cfg.Port(), a.initRoutes())
}
