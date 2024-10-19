package shortenerapi

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Makovey/shortener/internal/api"
	"github.com/Makovey/shortener/internal/config"
	"github.com/Makovey/shortener/internal/logger"
	"github.com/Makovey/shortener/internal/logger/stdout"
	"github.com/Makovey/shortener/internal/service/shortener"
)

type errReader int

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("error")
}

func TestPostNewURLHandler(t *testing.T) {
	type dependencies struct {
		service api.Shortener
		logger  logger.Logger
		config  config.HTTPConfig
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
				logger:  stdout.NewLoggerStdoutMock(),
				config:  config.NewHTTPConfig(),
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
				logger:  stdout.NewLoggerStdoutMock(),
				config:  config.NewHTTPConfig(),
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
				logger:  stdout.NewLoggerStdoutMock(),
				config:  config.NewHTTPConfig(),
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
			h := NewShortenerHandler(tt.dependencies.service, tt.dependencies.logger, tt.dependencies.config)
			r := httptest.NewRequest(http.MethodPost, "/", tt.parameters.body)
			w := httptest.NewRecorder()
			h.PostNewURLHandler(w, r)

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

func Test_handler_GetURLHandler(t *testing.T) {
	type dependencies struct {
		service api.Shortener
		logger  logger.Logger
		config  config.HTTPConfig
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
				logger:  stdout.NewLoggerStdoutMock(),
				config:  config.NewHTTPConfig(),
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
				logger:  stdout.NewLoggerStdoutMock(),
				config:  config.NewHTTPConfig(),
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
				logger:  stdout.NewLoggerStdoutMock(),
				config:  config.NewHTTPConfig(),
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
			h := NewShortenerHandler(tt.dependencies.service, tt.dependencies.logger, tt.dependencies.config)

			r := httptest.NewRequest(http.MethodGet, "/", nil)
			r.SetPathValue("id", tt.parameters.pathValue)

			w := httptest.NewRecorder()
			h.GetURLHandler(w, r)

			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, tt.want.code, res.StatusCode)
			assert.Equal(t, tt.want.location, res.Header.Get("Location"))
		})
	}
}
