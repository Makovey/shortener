package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/Makovey/shortener/internal/repository"
	"github.com/Makovey/shortener/internal/transport/model"
)

func (h handler) PostShortenURL(w http.ResponseWriter, r *http.Request) {
	fn := "http.PostShortenURL"

	userID := getUserIDFromContext(r.Context())
	if userID == "" || len(userID) != uuidLength {
		h.writeResponseWithError(w, http.StatusBadRequest, reloginAndTryAgain)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Error(fmt.Sprintf("[%s]: %s", fn, err.Error()))
		h.writeResponseWithError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	var req model.ShortenRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		h.logger.Error(fmt.Sprintf("[%s]: %s", fn, err.Error()))
		h.writeResponseWithError(w, http.StatusBadRequest, "body is invalid")
		return
	}

	if len(req.URL) == 0 {
		h.logger.Error(fmt.Sprintf("[%s]: can't short url, request body is empty", fn))
		h.writeResponseWithError(w, http.StatusBadRequest, "request body is empty")
		return
	}

	short, err := h.service.CreateShortURL(r.Context(), req.URL, userID)
	if err != nil {
		h.logger.Error(err.Error())
		if errors.Is(err, repository.ErrURLIsAlreadyExists) {
			h.writeResponse(w, http.StatusConflict, model.ShortenResponse{Result: short})
			return
		}
		h.writeResponseWithError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	h.writeResponse(w, http.StatusCreated, model.ShortenResponse{Result: short})
}
