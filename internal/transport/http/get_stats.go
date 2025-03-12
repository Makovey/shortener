package http

import (
	"fmt"
	"net/http"
)

// GetStats хендлер /api/internal/stats
// Отдает статистику по сервису
func (h handler) GetStats(w http.ResponseWriter, r *http.Request) {
	fn := "http.GetStats"

	userID := getUserIDFromContext(r.Context())
	if userID == "" || len(userID) != uuidLength {
		h.writeResponseWithError(w, http.StatusBadRequest, reloginAndTryAgain)
		return
	}

	model, err := h.service.GetStats(r.Context())
	if err != nil {
		h.logger.Error(fmt.Sprintf("[%s]: %v", fn, err))
		h.writeResponseWithError(w, http.StatusBadRequest, "internal server error")
		return
	}

	h.writeResponse(w, http.StatusOK, model)
}
