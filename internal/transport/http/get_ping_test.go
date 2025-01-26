package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Makovey/shortener/internal/logger/stdout"
	"github.com/Makovey/shortener/internal/service/mock"
	"github.com/Makovey/shortener/internal/service/shortener"
)

func TestGetPing(t *testing.T) {
	type dependencies struct {
		checker Checker
	}

	type want struct {
		code int
	}

	tests := []struct {
		name         string
		dependencies dependencies
		want         want
	}{
		{
			name: "successful pinged repo",
			dependencies: dependencies{
				checker: mock.NewCheckerMock(nil),
			},
			want: want{
				code: http.StatusOK,
			},
		},
		{
			name: "failed to ping repo: checker error",
			dependencies: dependencies{
				checker: mock.NewCheckerMock(errors.New("test error")),
			},
			want: want{
				code: http.StatusInternalServerError,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHTTPHandler(
				shortener.NewMockService(nil, nil),
				stdout.NewLoggerDummy(),
				tt.dependencies.checker,
			)

			r := httptest.NewRequest(http.MethodGet, "/ping", nil)
			w := httptest.NewRecorder()
			h.GetPing(w, r)

			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, tt.want.code, res.StatusCode)
		})
	}
}
