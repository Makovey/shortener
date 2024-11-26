package shortenerapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	def "github.com/Makovey/shortener/internal/api"
	"github.com/Makovey/shortener/internal/api/model"
	"github.com/Makovey/shortener/internal/logger"
	"github.com/Makovey/shortener/internal/middleware"
	"github.com/Makovey/shortener/internal/repository"
)

type handler struct {
	service def.Shortener
	checker def.Checker
	logger  logger.Logger
}

func (h handler) PostNewURL(w http.ResponseWriter, r *http.Request) {
	userID := getUserIDFromContext(r.Context())
	if userID == "" {
		h.writeResponseWithError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

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

	short, err := h.service.Short(string(longURL), userID)
	if err != nil {
		h.logger.Error(err.Error())
		if errors.Is(err, repository.ErrURLIsAlreadyExists) {
			w.WriteHeader(http.StatusConflict)
			_, _ = w.Write([]byte(short))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte(short))
}

func (h handler) GetURL(w http.ResponseWriter, r *http.Request) {
	userID := getUserIDFromContext(r.Context())
	if userID == "" {
		h.writeResponseWithError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	shortURL := r.PathValue("id")
	if len(shortURL) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	m, err := h.service.Get(shortURL, userID)
	if err != nil {
		h.logger.Error(err.Error())
		h.writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if m.IsDeleted {
		w.WriteHeader(http.StatusGone)
		return
	}

	w.Header().Add("Location", m.OriginalURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (h handler) PostShortenURL(w http.ResponseWriter, r *http.Request) {
	userID := getUserIDFromContext(r.Context())
	if userID == "" {
		h.writeResponseWithError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

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

	short, err := h.service.Short(req.URL, userID)
	if err != nil {
		h.logger.Error(err.Error())
		if errors.Is(err, repository.ErrURLIsAlreadyExists) {
			h.writeResponse(w, http.StatusConflict, model.ShortenResponse{Result: short})
			return
		}
		h.writeResponseWithError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	h.logger.Info(fmt.Sprintf("new short url created: %s", short))
	h.writeResponse(w, http.StatusCreated, model.ShortenResponse{Result: short})
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

func (h handler) PostBatch(w http.ResponseWriter, r *http.Request) {
	userID := getUserIDFromContext(r.Context())
	if userID == "" {
		h.writeResponseWithError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Error(fmt.Sprintf("can't read request body, cause: %s", err.Error()))
		h.writeResponseWithError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	var req []model.ShortenBatchRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		h.logger.Error(fmt.Sprintf("can't unmarshal request body, cause: %s", err.Error()))
		h.writeResponseWithError(w, http.StatusBadRequest, "body is invalid")
		return
	}

	if len(req) == 0 {
		h.logger.Error("can't short url, request body is empty")
		h.writeResponseWithError(w, http.StatusBadRequest, "request body is empty")
		return
	}

	resp, err := h.service.ShortBatch(req, userID)
	if err != nil {
		h.logger.Error(err.Error())
		h.writeResponseWithError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	h.logger.Info(fmt.Sprintf("batch processed with response: %s", resp))
	h.writeResponse(w, http.StatusCreated, resp)
}

func (h handler) GetAllURLS(w http.ResponseWriter, r *http.Request) {
	userID := getUserIDFromContext(r.Context())
	if userID == "" {
		h.writeResponseWithError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	models, err := h.service.GetAll(userID)
	if err != nil {
		h.logger.Error(err.Error())
		h.writeResponseWithError(w, http.StatusBadRequest, "internal server error")
		return
	}

	h.logger.Info(fmt.Sprintf("get all processed queried successfully with length: %d", len(models)))
	if len(models) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	h.writeResponse(w, http.StatusOK, models)
}

func (h handler) DeleteURLS(w http.ResponseWriter, r *http.Request) {
	userID := getUserIDFromContext(r.Context())
	if userID == "" {
		h.writeResponseWithError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Error(fmt.Sprintf("can't read request body, cause: %s", err.Error()))
		h.writeResponseWithError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	var ids []string
	err = json.Unmarshal(body, &ids)
	if err != nil {
		h.logger.Error(fmt.Sprintf("can't unmarshal request body, cause: %s", err.Error()))
		h.writeResponseWithError(w, http.StatusBadRequest, "body is invalid")
		return
	}

	deleteErrors := h.service.DeleteUsersURLS(r.Context(), userID, ids)

	for _, err = range deleteErrors {
		if err != nil {
			h.logger.Error(err.Error())
		}
	}

	w.WriteHeader(http.StatusAccepted)
}

func getUserIDFromContext(ctx context.Context) string {
	if ctx.Value(middleware.CtxUserIDKey) == nil {
		return ""
	}

	return ctx.Value(middleware.CtxUserIDKey).(string)
}

func NewShortenerHandler(
	service def.Shortener,
	logger logger.Logger,
	checker def.Checker,
) def.HTTPHandler {
	return &handler{
		service: service,
		logger:  logger,
		checker: checker,
	}
}
