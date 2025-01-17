package http

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Makovey/shortener/internal/config"
	"github.com/Makovey/shortener/internal/logger"
	"github.com/Makovey/shortener/internal/logger/stdout"
	"github.com/Makovey/shortener/internal/middleware"
	"github.com/Makovey/shortener/internal/service/dummy"
	"github.com/Makovey/shortener/internal/service/shortener"
)

func TestPostNewURLHandler(t *testing.T) {
	type dependencies struct {
		service Service
		logger  logger.Logger
		config  config.Config
		checker Checker
	}

	type want struct {
		code      int
		emptyBody bool
	}

	type parameters struct {
		body io.Reader
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
				service: shortener.NewMockService(false),
				logger:  stdout.NewLoggerDummy(),
				config:  config.NewConfig(stdout.NewLoggerDummy()),
				checker: dummy.NewDummyChecker(),
			},
			parameters: parameters{
				body: strings.NewReader("https://github.com"),
			},
			want: want{
				code:      http.StatusCreated,
				emptyBody: false,
			},
		},
		{
			name: "failed post new url: empty body",
			dependencies: dependencies{
				service: shortener.NewMockService(false),
				logger:  stdout.NewLoggerDummy(),
				config:  config.NewConfig(stdout.NewLoggerDummy()),
				checker: dummy.NewDummyChecker(),
			},
			parameters: parameters{
				body: strings.NewReader(""),
			},
			want: want{
				code:      http.StatusBadRequest,
				emptyBody: true,
			},
		},
		{
			name: "failed post new url: error with reader",
			dependencies: dependencies{
				service: shortener.NewMockService(true),
				logger:  stdout.NewLoggerDummy(),
				config:  config.NewConfig(stdout.NewLoggerDummy()),
				checker: dummy.NewDummyChecker(),
			},
			parameters: parameters{
				body: errReader(0),
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
				tt.dependencies.logger,
				tt.dependencies.checker,
			)
			r := httptest.NewRequest(http.MethodPost, "/", tt.parameters.body)
			w := httptest.NewRecorder()
			ctx := context.WithValue(r.Context(), middleware.CtxUserIDKey, uuid.NewString())

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
