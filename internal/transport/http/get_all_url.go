package http

import (
	"fmt"
	"net/http"
)

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

	h.logger.Info(fmt.Sprintf("get all processed queried successfully with length: %d", len(models)))
	if len(models) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	h.writeResponse(w, http.StatusOK, models)
}
