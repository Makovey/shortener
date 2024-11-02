package shortenerapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	def "github.com/Makovey/shortener/internal/api"
	"github.com/Makovey/shortener/internal/api/model"
	"github.com/Makovey/shortener/internal/config"
	"github.com/Makovey/shortener/internal/logger"
)

type handler struct {
	service def.Shortener
	checker def.Checker
	logger  logger.Logger
	config  config.Config
}

func (h handler) PostNewURLHandler(w http.ResponseWriter, r *http.Request) {
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
		h.logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (h handler) GetURLHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := r.PathValue("id")
	if len(shortURL) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	longURL, err := h.service.Get(shortURL)
	if err != nil {
		h.logger.Error(err.Error())
		h.writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Add("Location", longURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (h handler) PostShortenURLHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Error(fmt.Sprintf("can't read request body, cause: %s", err.Error()))
		h.writeResponseWithError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	var req model.ShortenRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		h.logger.Error(fmt.Sprintf("can't unmarshal request body, cause: %s", err.Error()))
		h.writeResponseWithError(w, http.StatusBadRequest, "body is invalid")
		return
	}

	if len(req.URL) == 0 {
		h.logger.Error("can't short url, request body is empty")
		h.writeResponseWithError(w, http.StatusBadRequest, "request body is empty")
		return
	}

	short := h.service.Short(req.URL)

	h.logger.Info(fmt.Sprintf("new short url created: %s", short))
	h.writeResponse(w, http.StatusCreated, model.ShortenResponse{Result: fmt.Sprintf("%s/%s", h.config.BaseReturnedURL(), short)})
}

func (h handler) GetPing(w http.ResponseWriter, r *http.Request) {
	err := h.checker.CheckPing()
	if err != nil {
		h.logger.Error(fmt.Sprintf("Ping error: %s", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func NewShortenerHandler(
	service def.Shortener,
	logger logger.Logger,
	config config.Config,
	checker def.Checker,
) def.HTTPHandler {
	return &handler{
		service: service,
		logger:  logger,
		config:  config,
		checker: checker,
	}
}
