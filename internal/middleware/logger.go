package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Makovey/shortener/internal/logger"
)

// Logger структура для содержания методов middleware
type Logger struct {
	log logger.Logger
}

// NewLogger конструктор Logger
func NewLogger(log logger.Logger) Logger {
	return Logger{log: log}
}

// Logger метод который выступает middleware, логирует статусы http ответов
func (l Logger) Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now()
		ww := wrappedResponseWriter{ResponseWriter: w}
		scheme := "http"
		if r.TLS != nil {
			scheme = "https"
		}

		defer func() {
			rInfo := fmt.Sprintf("%s %s://%s%s", r.Method, scheme, r.Host, r.RequestURI)
			l.log.Info(rInfo + fmt.Sprintf(" - %d %dB in %s", ww.code, ww.bodySize, time.Since(tm)))
		}()

		next.ServeHTTP(&ww, r)
	})
}

type wrappedResponseWriter struct {
	http.ResponseWriter
	wroteHeader bool
	code        int
	bodySize    int
}

// WriteHeader для соответствия интерфейсу
// Дополнительный флаг, для избежания коллизиий
func (b *wrappedResponseWriter) WriteHeader(code int) {
	if !b.wroteHeader {
		b.code = code
		b.wroteHeader = true
		b.ResponseWriter.WriteHeader(code)
	}
}

// Write дополнительно записывает размер ответа
func (b *wrappedResponseWriter) Write(buf []byte) (n int, err error) {
	b.writeHeaderIfNeeded()
	n, err = b.ResponseWriter.Write(buf)
	b.bodySize += n
	return n, err
}

func (b *wrappedResponseWriter) writeHeaderIfNeeded() {
	if !b.wroteHeader {
		b.WriteHeader(http.StatusOK)
	}
}
