package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/Makovey/shortener/internal/config"
	"github.com/Makovey/shortener/internal/logger"
	"github.com/Makovey/shortener/internal/logger/stdout"
	"github.com/Makovey/shortener/internal/middleware"
	"github.com/Makovey/shortener/internal/service/dummy"
	"github.com/Makovey/shortener/internal/service/shortener"
)

func TestGetURLHandler(t *testing.T) {
	type dependencies struct {
		service Service
		logger  logger.Logger
		config  config.Config
		checker Checker
	}

	type want struct {
		code     int
		location string
	}

	type parameters struct {
		pathValue string
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
				service: shortener.NewMockService(false),
				logger:  stdout.NewLoggerDummy(),
				config:  config.NewConfig(stdout.NewLoggerDummy()),
				checker: dummy.NewDummyChecker(),
			},
			parameters: parameters{
				pathValue: "/a1b2c3",
			},
			want: want{
				code:     http.StatusTemporaryRedirect,
				location: "https://github.com",
			},
		},
		{
			name: "failed get long url: empty path value",
			dependencies: dependencies{
				service: shortener.NewMockService(false),
				logger:  stdout.NewLoggerDummy(),
				config:  config.NewConfig(stdout.NewLoggerDummy()),
				checker: dummy.NewDummyChecker(),
			},
			parameters: parameters{
				pathValue: "",
			},
			want: want{
				code:     http.StatusBadRequest,
				location: "",
			},
		},
		{
			name: "failed get long url: error from service",
			dependencies: dependencies{
				service: shortener.NewMockService(true),
				logger:  stdout.NewLoggerDummy(),
				config:  config.NewConfig(stdout.NewLoggerDummy()),
				checker: dummy.NewDummyChecker(),
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
				tt.dependencies.logger,
				tt.dependencies.checker,
			)

			r := httptest.NewRequest(http.MethodGet, "/", nil)
			r.SetPathValue("id", tt.parameters.pathValue)
			ctx := context.WithValue(r.Context(), middleware.CtxUserIDKey, uuid.NewString())

			w := httptest.NewRecorder()
			h.GetURL(w, r.WithContext(ctx))

			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, tt.want.code, res.StatusCode)
			assert.Equal(t, tt.want.location, res.Header.Get("Location"))
		})
	}
}
