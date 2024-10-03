package app

import "net/http"

type App struct {
	dependencyProvider *dependencyProvider
}

func NewApp() *App {
	return &App{dependencyProvider: newDependencyProvider()}
}

func (a *App) Run() error {
	return a.runHttpServer()
}

func (a *App) initRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /", a.dependencyProvider.HttpHandler().PostNewUrlHandler)
	mux.HandleFunc("GET /{id}", a.dependencyProvider.HttpHandler().GetUrlHandler)

	return mux
}

func (a *App) runHttpServer() error {
	cfg := a.dependencyProvider.Config()

	return http.ListenAndServe(cfg.Port(), a.initRoutes())
}
