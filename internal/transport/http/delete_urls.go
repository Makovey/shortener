package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// DeleteURLS хендлер /api/user/urls
// ждет на вход список коротких урлов для их дальнейшего удаления
func (h handler) DeleteURLS(w http.ResponseWriter, r *http.Request) {
	fn := "http.DeleteURLS"

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

	var ids []string
	err = json.Unmarshal(body, &ids)
	if err != nil {
		h.logger.Error(fmt.Sprintf("[%s]: %s", fn, err.Error()))
		h.writeResponseWithError(w, http.StatusBadRequest, "body is invalid")
		return
	}

	if len(ids) == 0 {
		h.logger.Error(fmt.Sprintf("[%s]: ids is empty", fn))
		h.writeResponseWithError(w, http.StatusBadRequest, "body is empty")
		return
	}

	deleteErrors := h.service.DeleteUsersURLs(r.Context(), userID, ids)

	for _, err = range deleteErrors {
		if err != nil {
			h.logger.Error(fmt.Sprintf("[%s]: %s", fn, err.Error()))
		}
	}

	w.WriteHeader(http.StatusAccepted)
}
