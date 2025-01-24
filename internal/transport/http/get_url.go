package http

import (
	"fmt"
	"net/http"
)

func (h handler) GetURL(w http.ResponseWriter, r *http.Request) {
	fn := "http.GetURL"

	userID := getUserIDFromContext(r.Context())
	if userID == "" || len(userID) != uuidLength {
		h.writeResponseWithError(w, http.StatusBadRequest, reloginAndTryAgain)
		return
	}

	shortURL := r.PathValue("id")
	if len(shortURL) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	url, err := h.service.GetFullURL(r.Context(), shortURL, userID)
	if err != nil {
		h.logger.Error(fmt.Sprintf("[%s]: %s", fn, err.Error()))
		h.writeResponse(w, http.StatusBadRequest, "bad request")
		return
	}

	if url.IsDeleted {
		w.WriteHeader(http.StatusGone)
		return
	}

	w.Header().Add("Location", url.OriginalURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
