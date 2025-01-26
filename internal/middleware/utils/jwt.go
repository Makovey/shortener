package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/Makovey/shortener/internal/logger"
)

const (
	tokenExp  = time.Hour * 24
	secretKey = "shortener_secret_key"
)

// Ошибки, которые может вернуть валидатор JWT токенов
var (
	ErrSigningMethod = errors.New("unexpected signing method")
	ErrParseToken    = errors.New("failed to parse token")
	ErrInvalidToken  = errors.New("invalid token")
	ErrTokenExpired  = errors.New("token is expired")
)

// JWTUtils хелперы по работе с JWT
type JWTUtils struct {
	logger logger.Logger
}

// NewJWTUtils конструктор JWTUtils
func NewJWTUtils(logger logger.Logger) JWTUtils {
	return JWTUtils{logger: logger}
}

// Claims полезная информация в токене
type Claims struct {
	jwt.RegisteredClaims
	UserID string
}

// BuildNewJWT генерирует новый JWT токен, которы действует tokenExp
func (j JWTUtils) BuildNewJWT(userID string) (string, error) {
	f := "jwt.buildNewJWT:"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExp)),
		},
		UserID: userID,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv(secretKey)))
	if err != nil {
		j.logger.Warning(fmt.Sprintf("%s can't sign token: %s", f, err.Error()))
		return "", err
	}

	return tokenString, nil
}

// ParseUserID принимает JWT токен, и возвращает UserID, если токен валидный и действующий
func (j JWTUtils) ParseUserID(tokenString string) (string, error) {
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
