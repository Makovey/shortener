package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"os"
	"time"

	"github.com/Makovey/shortener/internal/logger"
	"github.com/golang-jwt/jwt/v5"
)

type Key string

const (
	tokenExp          = time.Hour * 24
	jwtCookieName     = "jwt"
	secretKey         = "JWT_SECRET" // TODO: replace to env
	CtxUserIDKey  Key = "UserID"
)

var (
	ErrSigningMethod = errors.New("unexpected signing method")
	ErrParseToken    = errors.New("failed to parse token")
	ErrInvalidToken  = errors.New("invalid token")
	ErrTokenExpired  = errors.New("token is expired")
)

type JWTHandler struct {
	logger logger.Logger
}

type Claims struct {
	jwt.RegisteredClaims
	UserID string
}

func (j JWTHandler) JWTHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f := "jwt.JWTHandler:"
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

		if isCookieAbsent && r.URL.Path == "/api/user/urls" {
			responseWithError(w, http.StatusUnauthorized, "unauthorized")
			return
		}

		if !isCookieAbsent {
			userID, err = j.parseUserID(jwtCookie.Value)
			if err != nil && errors.Is(err, ErrParseToken) {
				responseWithError(w, http.StatusInternalServerError, "internal server error")
				return
			}

			if userID == "" {
				j.logger.Warning(fmt.Sprintf("%s UserID is empty", f))
				responseWithError(w, http.StatusUnauthorized, "unauthorized")
				return
			}
		}

		if isCookieAbsent || errors.Is(err, ErrTokenExpired) || errors.Is(err, ErrInvalidToken) {
			userID = uuid.NewString()[:6]
			tokenString, err := j.buildNewJWT(userID)
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

func (j JWTHandler) buildNewJWT(userID string) (string, error) {
	f := "jwt.buildNewJWT:"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExp)),
		},
		UserID: userID,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv(secretKey)))
	if err != nil {
		j.logger.Warning(fmt.Sprintf("%s Can't sign token: %s", f, err.Error()))
		return "", err
	}

	return tokenString, nil
}

func (j JWTHandler) parseUserID(tokenString string) (string, error) {
	f := "jwt.parseUserID:"
	var claims Claims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			j.logger.Warning(fmt.Sprintf("%s unexpected signing method: %v", f, token.Header["alg"]))
			return nil, ErrSigningMethod
		}

		return []byte(os.Getenv(secretKey)), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return "", ErrTokenExpired
		}

		j.logger.Warning(fmt.Sprintf("%s failed to parse token -> %s", f, err.Error()))
		return "", ErrParseToken
	}

	if !token.Valid {
		j.logger.Warning(fmt.Sprintf("%s token is invalid", f))
		return "", ErrInvalidToken
	}

	return claims.UserID, nil
}

func NewJWTHandler(log logger.Logger) JWTHandler {
	return JWTHandler{logger: log}
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
