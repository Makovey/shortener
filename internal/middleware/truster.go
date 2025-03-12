package middleware

import (
	"net"
	"net/http"
)

// Truster проверяет, доступно ли пользователю хендлеры
type Truster struct {
	trustedSubnet string
}

// NewTruster конструктор для middleware
func NewTruster(
	trustedSubnet string,
) *Truster {
	return &Truster{
		trustedSubnet: trustedSubnet,
	}
}

// CheckSubnet проверяет, принадлежность хоста к разрешённой подсети
func (t *Truster) CheckSubnet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.Header.Get("X-Real-IP")
		if ip == "" {
			responseWithError(w, http.StatusForbidden, "ip address is required")
			return
		}

		_, trustedIP, err := net.ParseCIDR(t.trustedSubnet)
		if err != nil {
			responseWithError(w, http.StatusBadRequest, "invalid trusted subnet")
			return
		}

		if trustedIP.Contains(net.ParseIP(ip)) {
			responseWithError(w, http.StatusForbidden, "your ip address not trusted")
			return
		}

		next.ServeHTTP(w, r)
	})
}
