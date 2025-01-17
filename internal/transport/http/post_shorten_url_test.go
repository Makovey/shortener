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

	"github.com/Makovey/shortener/internal/config"
	"github.com/Makovey/shortener/internal/logger"
	"github.com/Makovey/shortener/internal/logger/stdout"
	"github.com/Makovey/shortener/internal/middleware"
	"github.com/Makovey/shortener/internal/service/dummy"
	"github.com/Makovey/shortener/internal/service/shortener"
)

func TestPostShortenHandler(t *testing.T) {
	type dependencies struct {
		service Service
		logger  logger.Logger
		config  config.Config
		checker Checker
	}

	type want struct {
		code         int
		containsBody string
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
			name: "successful post transport shorten",
			dependencies: dependencies{
				service: shortener.NewMockService(false),
				logger:  stdout.NewLoggerDummy(),
				config:  config.NewConfig(stdout.NewLoggerDummy()),
				checker: dummy.NewDummyChecker(),
			},
			parameters: parameters{
				body: strings.NewReader(makeJSON(map[string]any{
					"url": "https://github.com",
				})),
			},
			want: want{
				code:         http.StatusCreated,
				containsBody: "result",
			},
		},
		{
			name: "successful post transport shorten",
			dependencies: dependencies{
				service: shortener.NewMockService(false),
				logger:  stdout.NewLoggerDummy(),
				config:  config.NewConfig(stdout.NewLoggerDummy()),
				checker: dummy.NewDummyChecker(),
			},
			parameters: parameters{
				body: strings.NewReader(makeJSON(map[string]any{
					"url": "https://github.com",
				})),
			},
			want: want{
				code:         http.StatusCreated,
				containsBody: "result",
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
				code:         http.StatusInternalServerError,
				containsBody: "error",
			},
		},
		{
			name: "failed post new url: invalid body",
			dependencies: dependencies{
				service: shortener.NewMockService(false),
				logger:  stdout.NewLoggerDummy(),
				config:  config.NewConfig(stdout.NewLoggerDummy()),
				checker: dummy.NewDummyChecker(),
			},
			parameters: parameters{
				body: strings.NewReader(makeJSON(map[string]any{
					"url": 1234567890,
				})),
			},
			want: want{
				code:         http.StatusBadRequest,
				containsBody: "error",
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
				body: strings.NewReader(makeJSON(map[string]any{
					"url": "",
				})),
			},
			want: want{
				code:         http.StatusBadRequest,
				containsBody: "error",
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

			r := httptest.NewRequest(http.MethodPost, "/transport/shorten", tt.parameters.body)
			ctx := context.WithValue(r.Context(), middleware.CtxUserIDKey, uuid.NewString())

			w := httptest.NewRecorder()
			h.PostShortenURL(w, r.WithContext(ctx))

			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, tt.want.code, res.StatusCode)
			assert.Contains(t, parseBody(res.Body), tt.want.containsBody)
		})
	}
}
