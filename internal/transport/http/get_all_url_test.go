package http

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	comModel "github.com/Makovey/shortener/internal/service/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/Makovey/shortener/internal/logger/stdout"
	"github.com/Makovey/shortener/internal/middleware"
	"github.com/Makovey/shortener/internal/service/mock"
	"github.com/Makovey/shortener/internal/service/shortener"
)

func TestGetAllURLSHandler(t *testing.T) {
	type dependencies struct {
		service Service
	}

	type want struct {
		code int
	}

	type parameters struct {
		userID string
	}

	tests := []struct {
		name         string
		dependencies dependencies
		parameters   parameters
		want         want
	}{
		{
			name: "successful get all urls",
			dependencies: dependencies{
				service: shortener.NewMockService(
					nil,
					[]comModel.ShortenBatch{
						{
							CorrelationID: "1",
						},
					},
				),
			},
			parameters: parameters{
				userID: uuid.NewString(),
			},
			want: want{
				code: http.StatusOK,
			},
		},
		{
			name: "successful get all urls, but service returned empty list",
			dependencies: dependencies{
				service: shortener.NewMockService(
					nil,
					[]comModel.ShortenBatch{},
				),
			},
			parameters: parameters{
				userID: uuid.NewString(),
			},
			want: want{
				code: http.StatusNoContent,
			},
		},
		{
			name: "failed get all urls: user id is empty",
			dependencies: dependencies{
				service: shortener.NewMockService(nil, nil),
			},
			parameters: parameters{},
			want: want{
				code: http.StatusBadRequest,
			},
		},
		{
			name: "failed get long url: error from service",
			dependencies: dependencies{
				service: shortener.NewMockService(errors.New("test error"), nil),
			},
			parameters: parameters{
				userID: uuid.NewString(),
			},
			want: want{
				code: http.StatusBadRequest,
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

			r := httptest.NewRequest(http.MethodGet, "/api/user/urls", nil)
			ctx := context.WithValue(r.Context(), middleware.CtxUserIDKey, tt.parameters.userID)

			w := httptest.NewRecorder()
			h.GetAllURLS(w, r.WithContext(ctx))

			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, tt.want.code, res.StatusCode)
		})
	}
}
