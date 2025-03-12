package http

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/Makovey/shortener/internal/logger/stdout"
	"github.com/Makovey/shortener/internal/middleware"
	"github.com/Makovey/shortener/internal/service"
	"github.com/Makovey/shortener/internal/service/mock"
	"github.com/Makovey/shortener/internal/service/shortener"
	"github.com/Makovey/shortener/internal/transport/model"
)

func TestGetURLHandler(t *testing.T) {
	type dependencies struct {
		service service.Service
	}

	type want struct {
		code     int
		location string
	}

	type parameters struct {
		pathValue string
		userID    string
	}

	tests := []struct {
		name         string
		dependencies dependencies
		parameters   parameters
		want         want
	}{
		{
			name: "successful get url",
			dependencies: dependencies{
				service: shortener.NewMockService(nil, nil),
			},
			parameters: parameters{
				pathValue: "/a1b2c3",
				userID:    uuid.NewString(),
			},
			want: want{
				code:     http.StatusTemporaryRedirect,
				location: "https://github.com",
			},
		},
		{
			name: "successful get url, url already deleted",
			dependencies: dependencies{
				service: shortener.NewMockService(nil, model.UserFullURL{
					OriginalURL: "",
					IsDeleted:   true,
				}),
			},
			parameters: parameters{
				pathValue: "/a1b2c3",
				userID:    uuid.NewString(),
			},
			want: want{
				code:     http.StatusGone,
				location: "",
			},
		},
		{
			name: "failed get long url: empty path value",
			dependencies: dependencies{
				service: shortener.NewMockService(nil, nil),
			},
			parameters: parameters{
				pathValue: "",
				userID:    uuid.NewString(),
			},
			want: want{
				code:     http.StatusBadRequest,
				location: "",
			},
		},
		{
			name: "failed get long url: error from service",
			dependencies: dependencies{
				service: shortener.NewMockService(errors.New("test error"), nil),
			},
			parameters: parameters{
				pathValue: "/a1b2c3",
				userID:    uuid.NewString(),
			},
			want: want{
				code:     http.StatusBadRequest,
				location: "",
			},
		},
		{
			name: "failed get long url: user id is empty",
			dependencies: dependencies{
				service: shortener.NewMockService(nil, nil),
			},
			parameters: parameters{
				pathValue: "/a1b2c3",
			},
			want: want{
				code:     http.StatusBadRequest,
				location: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHTTPHandler(
				tt.dependencies.service,
				stdout.NewLoggerDummy(),
				mock.NewCheckerMock(nil),
			)

			r := httptest.NewRequest(http.MethodGet, "/", nil)
			r.SetPathValue("id", tt.parameters.pathValue)
			ctx := context.WithValue(r.Context(), middleware.CtxUserIDKey, tt.parameters.userID)

			w := httptest.NewRecorder()
			h.GetURL(w, r.WithContext(ctx))

			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, tt.want.code, res.StatusCode)
			assert.Equal(t, tt.want.location, res.Header.Get("Location"))
		})
	}
}
