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

	"github.com/Makovey/shortener/internal/logger/stdout"
	"github.com/Makovey/shortener/internal/middleware"
	"github.com/Makovey/shortener/internal/service/mock"
	"github.com/Makovey/shortener/internal/service/shortener"
)

func TestDeleteURLSHandler(t *testing.T) {
	type dependencies struct {
		service Service
	}

	type want struct {
		code int
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
			name: "successful post batch",
			dependencies: dependencies{
				service: shortener.NewMockService(nil, nil),
			},
			parameters: parameters{
				body:   strings.NewReader(makeList([]string{"1", "2"})),
				userID: uuid.NewString(),
			},
			want: want{
				code: http.StatusAccepted,
			},
		},
		{
			name: "failed to post batch: user id is empty",
			dependencies: dependencies{
				service: shortener.NewMockService(nil, nil),
			},
			parameters: parameters{
				body: strings.NewReader(makeList([]string{"1", "2"})),
			},
			want: want{
				code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to post batch: error with reader",
			dependencies: dependencies{
				service: shortener.NewMockService(nil, nil),
			},
			parameters: parameters{
				body:   errReader(0),
				userID: uuid.NewString(),
			},
			want: want{
				code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to post batch: error with empty body",
			dependencies: dependencies{
				service: shortener.NewMockService(nil, nil),
			},
			parameters: parameters{
				body:   strings.NewReader(""),
				userID: uuid.NewString(),
			},
			want: want{
				code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to post batch: empty list",
			dependencies: dependencies{
				service: shortener.NewMockService(nil, nil),
			},
			parameters: parameters{
				body:   strings.NewReader(makeList([]string{})),
				userID: uuid.NewString(),
			},
			want: want{
				code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to post batch: service error",
			dependencies: dependencies{
				service: shortener.NewMockService(errors.New("test error"), nil),
			},
			parameters: parameters{

				body:   strings.NewReader(makeList([]string{"1", "2"})),
				userID: uuid.NewString(),
			},
			want: want{
				code: http.StatusAccepted,
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

			r := httptest.NewRequest(http.MethodDelete, "/api/user/urls", tt.parameters.body)
			ctx := context.WithValue(r.Context(), middleware.CtxUserIDKey, tt.parameters.userID)

			w := httptest.NewRecorder()
			h.DeleteURLS(w, r.WithContext(ctx))

			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, tt.want.code, res.StatusCode)
		})
	}
}
