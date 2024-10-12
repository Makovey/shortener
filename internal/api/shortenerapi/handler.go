package shortenerapi

import (
	"fmt"
	"io"
	"net/http"

	def "github.com/Makovey/shortener/internal/api"
	"github.com/Makovey/shortener/internal/config"
	"github.com/Makovey/shortener/internal/logger"
)

type handler struct {
	service def.Shortener
	logger  logger.Logger
	config  config.HTTPConfig
}

func (h *handler) PostNewURLHandler(w http.ResponseWriter, r *http.Request) {
	longURL, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Error(fmt.Sprintf("Can't read request body, cause: %s", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(longURL) == 0 {
		h.logger.Error("Can't short url, cause request body is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	short := h.service.Short(string(longURL))
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte(fmt.Sprintf("%s/%s", h.config.BaseReturnedURL(), short)))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (h *handler) GetURLHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := r.PathValue("id")
	if shortURL == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	longURL, err := h.service.Get(shortURL)
	if err != nil {
		h.logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Location", longURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func NewShortenerHandler(
	service def.Shortener, logger logger.Logger, config config.HTTPConfig) def.HTTPHandler {
	return &handler{
		service: service,
		logger:  logger,
		config:  config,
	}
}
