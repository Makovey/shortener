package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"github.com/Makovey/shortener/internal/logger"
	"github.com/Makovey/shortener/internal/middleware/utils"
)

// Key отдельный тип для ключа используемый контекстом, во избежание коллизий
type Key string

// Ключи, по которым лежат авторизационные данные
const (
	jwtCookieName     = "jwt"
	CtxUserIDKey  Key = "UserID"
)

// AuthHandler отвечает за авторизацию в приложении
type AuthHandler struct {
	jwtUtils utils.JWTUtils
	logger   logger.Logger
}

// NewAuthHandler конструктор для AuthHandler
func NewAuthHandler(
	log logger.Logger,
	jwtUtils utils.JWTUtils,
) AuthHandler {
	return AuthHandler{logger: log, jwtUtils: jwtUtils}
}

// AuthHandler выступает middleware авторизации
// Проверяет в куках наличие JWT токена и проверяет его на валидность
func (j AuthHandler) AuthHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f := "jwt.AuthHandler:"
		var isCookieAbsent bool
		var userID string

		jwtCookie, err := r.Cookie(jwtCookieName)
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				isCookieAbsent = true
			} else {
				responseWithError(w, http.StatusInternalServerError, "internal server error")
				return
			}
		}

		if !isCookieAbsent {
			userID, err = j.jwtUtils.ParseUserID(jwtCookie.Value)
			if err != nil && errors.Is(err, utils.ErrParseToken) {
				responseWithError(w, http.StatusInternalServerError, "internal server error")
				return
			}

			if userID == "" {
				j.logger.Warning(fmt.Sprintf("%s UserID is empty", f))
				responseWithError(w, http.StatusUnauthorized, "unauthorized")
				return
			}
		}

		if isCookieAbsent || errors.Is(err, utils.ErrTokenExpired) || errors.Is(err, utils.ErrInvalidToken) {
			userID = uuid.NewString()
			tokenString, err := j.jwtUtils.BuildNewJWT(userID)
			if err != nil {
				responseWithError(w, http.StatusInternalServerError, "internal server error")
				return
			}

			http.SetCookie(w, &http.Cookie{
				Name:     jwtCookieName,
				Value:    tokenString,
				HttpOnly: true,
			})
		}

		ctx := context.WithValue(r.Context(), CtxUserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func responseWithError(w http.ResponseWriter, status int, message string) {
	type Response struct {
		Error string `json:"error"`
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response := Response{
		Error: message,
	}

	json.NewEncoder(w).Encode(response)
}
