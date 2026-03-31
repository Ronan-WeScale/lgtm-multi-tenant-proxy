package httputil

import (
	"net/http"
	"strings"
)

// RealClientIP returns the real client IP from X-Real-Ip or X-Forwarded-For headers,
// falling back to RemoteAddr.
func RealClientIP(r *http.Request) string {
	if ip := r.Header.Get("X-Real-Ip"); ip != "" {
		return ip
	}
	if ip := r.Header.Get("X-Forwarded-For"); ip != "" {
		return strings.SplitN(ip, ",", 2)[0]
	}
	return r.RemoteAddr
}
