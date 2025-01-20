package http

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/Makovey/shortener/internal/repository"
)

func (h handler) PostNewURL(w http.ResponseWriter, r *http.Request) {
	fn := "http.PostNewURL"

	userID := getUserIDFromContext(r.Context())
	if userID == "" || len(userID) != uuidLength {
		h.writeResponseWithError(w, http.StatusBadRequest, reloginAndTryAgain)
		return
	}

	longURL, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Error(fmt.Sprintf("[%s]: %s", fn, err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(longURL) == 0 {
		h.logger.Error(fmt.Sprintf("[%s]: can't short url, cause request body is empty", fn))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	short, err := h.service.CreateShortURL(r.Context(), string(longURL), userID)
	if err != nil {
		h.logger.Error(err.Error())
		if errors.Is(err, repository.ErrURLIsAlreadyExists) {
			w.WriteHeader(http.StatusConflict)
			_, err = w.Write([]byte(short))
			if err != nil {
				h.logger.Error(fmt.Sprintf("[%s]: %s", fn, err.Error()))
			}
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte(short))
	if err != nil {
		h.logger.Error(fmt.Sprintf("[%s]: %s", fn, err.Error()))
	}
}
