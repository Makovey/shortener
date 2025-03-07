package http

import (
	"context"
	"fmt"
	"net/http"
)

// Checker интерфейс по проверке компонентов
type Checker interface {
	CheckPing(ctx context.Context) error
}

// GetPing хендлер /ping
// Проверяет работу репозитория
func (h handler) GetPing(w http.ResponseWriter, r *http.Request) {
	fn := "http.GetPing"

	err := h.checker.CheckPing(r.Context())
	if err != nil {
		h.logger.Error(fmt.Sprintf("[%s]: %s", fn, err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
