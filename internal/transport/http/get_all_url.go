package http

import (
	"fmt"
	"net/http"
)

// GetAllURLS хендлер /api/user/urls
// Возвращает список урлов юзера
func (h handler) GetAllURLS(w http.ResponseWriter, r *http.Request) {
	fn := "http.GetAllURLS"

	userID := getUserIDFromContext(r.Context())
	if userID == "" || len(userID) != uuidLength {
		h.writeResponseWithError(w, http.StatusBadRequest, reloginAndTryAgain)
		return
	}

	models, err := h.service.GetAllURLs(r.Context(), userID)
	if err != nil {
		h.logger.Error(fmt.Sprintf("[%s]: %v", fn, err))
		h.writeResponseWithError(w, http.StatusBadRequest, "internal server error")
		return
	}

	if len(models) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	h.writeResponse(w, http.StatusOK, models)
}
