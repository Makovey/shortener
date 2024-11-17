package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/Makovey/shortener/internal/logger"
)

const (
	tokenExp      = time.Hour * 24
	jwtCookieName = "jwt"
	CtxUserIDKey  = "UserID"
)

var (
	ErrSigningMethod = errors.New("unexpected signing method")
	ErrParseToken    = errors.New("failed to parse token")
	ErrInvalidToken  = errors.New("invalid token")
)

type JWTHandler struct {
	logger logger.Logger
}

type Claims struct {
	jwt.RegisteredClaims
	UserID string
}

func (j JWTHandler) JWTHandler(next http.Handler) http.Handler {
	f := "jwt.JWTHandler:"
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwtCookie, err := r.Cookie(jwtCookieName)
		if err != nil {
			j.logger.Info(fmt.Sprintf("%s JWT not found", f))
			newUserID := uuid.NewString()[:6]
			tokenString, err := j.buildNewJWT(newUserID)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				next.ServeHTTP(w, r)
				return
			}

			http.SetCookie(w, &http.Cookie{
				Name:  jwtCookieName,
				Value: tokenString,
			})

			ctx := context.WithValue(r.Context(), CtxUserIDKey, newUserID)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			userID, err := j.parseUserID(jwtCookie.Value)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), CtxUserIDKey, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
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

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
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

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		j.logger.Warning(fmt.Sprintf("%s failed to parse token -> %s", f, err.Error()))
		return "", ErrParseToken
	}

	if !token.Valid {
		j.logger.Warning(fmt.Sprintf("%s token is invalid", f))
		return "", ErrInvalidToken
	}

	if claims.UserID == "" {
		j.logger.Warning(fmt.Sprintf("%s UserID is empty", f))
		return "", ErrInvalidToken
	}

	return claims.UserID, nil
}

func NewJWTHandler(log logger.Logger) JWTHandler {
	return JWTHandler{logger: log}
}
