package http

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Makovey/shortener/internal/logger/stdout"
	"github.com/Makovey/shortener/internal/middleware"
	"github.com/Makovey/shortener/internal/repository"
	"github.com/Makovey/shortener/internal/service/mock"
	"github.com/Makovey/shortener/internal/service/shortener"
)

func TestPostNewURLHandler(t *testing.T) {
	type dependencies struct {
		service Service
	}

	type want struct {
		code      int
		emptyBody bool
	}

	type parameters struct {
		body   io.Reader
		userID string
	}

	tests := []struct {
		name         string
		dependencies dependencies
		parameters   parameters
		want         want
	}{
		{
			name: "successful post new url",
			dependencies: dependencies{
				service: shortener.NewMockService(nil, nil),
			},
			parameters: parameters{
				body:   strings.NewReader("https://github.com"),
				userID: uuid.NewString(),
			},
			want: want{
				code:      http.StatusCreated,
				emptyBody: false,
			},
		},
		{
			name: "failed post new url: empty body",
			dependencies: dependencies{
				service: shortener.NewMockService(nil, nil),
			},
			parameters: parameters{
				body:   strings.NewReader(""),
				userID: uuid.NewString(),
			},
			want: want{
				code:      http.StatusBadRequest,
				emptyBody: true,
			},
		},
		{
			name: "failed post new url: error with reader",
			dependencies: dependencies{
				service: shortener.NewMockService(nil, nil),
			},
			parameters: parameters{
				body:   errReader(0),
				userID: uuid.NewString(),
			},
			want: want{
				code:      http.StatusBadRequest,
				emptyBody: true,
			},
		},
		{
			name: "failed post new url: service error",
			dependencies: dependencies{
				service: shortener.NewMockService(errors.New("test error"), nil),
			},
			parameters: parameters{
				body:   strings.NewReader("https://github.com"),
				userID: uuid.NewString(),
			},
			want: want{
				code:      http.StatusInternalServerError,
				emptyBody: true,
			},
		},
		{
			name: "failed post new url: service error, url already exists",
			dependencies: dependencies{
				service: shortener.NewMockService(repository.ErrURLIsAlreadyExists, nil),
			},
			parameters: parameters{
				body:   strings.NewReader("https://github.com"),
				userID: uuid.NewString(),
			},
			want: want{
				code:      http.StatusConflict,
				emptyBody: true,
			},
		},
		{
			name: "failed post new url: empty user id",
			dependencies: dependencies{
				service: shortener.NewMockService(repository.ErrURLIsAlreadyExists, nil),
			},
			parameters: parameters{
				body: strings.NewReader("https://github.com"),
			},
			want: want{
				code:      http.StatusBadRequest,
				emptyBody: true,
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
			r := httptest.NewRequest(http.MethodPost, "/", tt.parameters.body)
			w := httptest.NewRecorder()
			ctx := context.WithValue(r.Context(), middleware.CtxUserIDKey, tt.parameters.userID)

			h.PostNewURL(w, r.WithContext(ctx))

			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, tt.want.code, res.StatusCode)

			resBody, err := io.ReadAll(res.Body)
			require.NoError(t, err)

			if !tt.want.emptyBody {
				assert.NotEmpty(t, resBody)
			}
		})
	}
}
