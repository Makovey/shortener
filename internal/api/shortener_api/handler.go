package shortener_api

import (
	"fmt"
	"io"
	"net/http"

	def "github.com/Makovey/shortener/internal/api"
	"github.com/Makovey/shortener/internal/config"
	"github.com/Makovey/shortener/internal/logger"
	"github.com/Makovey/shortener/internal/service"
)

type handler struct {
	service service.ShortenerService
	logger  logger.Logger
	config  config.HttpConfig
}

func (h *handler) PostNewUrlHandler(w http.ResponseWriter, r *http.Request) {
	longUrl, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Error(fmt.Sprintf("Can't read request body, cause: %s", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(longUrl) == 0 {
		h.logger.Error("Can't short url, cause request body is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	short, err := h.service.Short(string(longUrl))
	if err != nil {
		h.logger.Error(fmt.Sprintf("Can't to short url, cause: %s", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte(fmt.Sprintf("http://localhost%s/%s", h.config.Port(), short)))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (h *handler) GetUrlHandler(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.PathValue("id")
	if shortUrl == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	longUrl, err := h.service.Get(shortUrl)
	if err != nil {
		h.logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Location", longUrl)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func NewShortenerHandler(
	service service.ShortenerService, logger logger.Logger, config config.HttpConfig) def.HttpHandler {
	return &handler{
		service: service,
		logger:  logger,
		config:  config,
	}
}
