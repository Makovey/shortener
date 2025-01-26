package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Makovey/shortener/internal/transport/model"
)

// PostBatch хендлер /api/shorten/batch
// Принимает список полных урлов, возвращает список коротких урлов
func (h handler) PostBatch(w http.ResponseWriter, r *http.Request) {
	fn := "http.PostBatch"

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

	var req []model.ShortenBatchRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		h.logger.Error(fmt.Sprintf("[%s]: %s", fn, err.Error()))
		h.writeResponseWithError(w, http.StatusBadRequest, "body is invalid")
		return
	}

	if len(req) == 0 {
		h.logger.Error(fmt.Sprintf("[%s]: can't short url, request body is empty", fn))
		h.writeResponseWithError(w, http.StatusBadRequest, "request body is empty")
		return
	}

	resp, err := h.service.ShortBatch(r.Context(), req, userID)
	if err != nil {
		h.logger.Error(fmt.Sprintf("[%s]: %s", fn, err.Error()))
		h.writeResponseWithError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	h.writeResponse(w, http.StatusCreated, resp)
}
