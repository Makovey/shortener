package app

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"

	"github.com/Makovey/shortener/internal/middleware"
	"github.com/Makovey/shortener/internal/middleware/utils"
)

func (a *App) initRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.NewLogger(a.log).Logger)

	jwtMiddleware := middleware.NewAuthHandler(a.log, utils.NewJWTUtils(a.log))
	r.Use(jwtMiddleware.AuthHandler)
	r.Use(middleware.NewCompressor().Compress)
	r.Use(chiMiddleware.Recoverer)

	r.Mount("/debug", chiMiddleware.Profiler())

	r.Post("/", a.handler.PostNewURL)
	r.Get("/{id}", a.handler.GetURL)
	r.Get("/ping", a.handler.GetPing)

	r.Route("/api", func(r chi.Router) {
		r.Post("/shorten", a.handler.PostShortenURL)
		r.Post("/shorten/batch", a.handler.PostBatch)
		r.Get("/user/urls", a.handler.GetAllURLS)
		r.Delete("/user/urls", a.handler.DeleteURLS)
	})

	r.Group(func(r chi.Router) {
		r.Use(middleware.NewTruster(a.cfg.TrustedSubnet()).CheckSubnet)
		r.Get("/api/internal/stats", a.handler.GetStats)
	})

	return r
}
