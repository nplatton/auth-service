package middleware

import (
	"net/http"
	"strings"
	"time"

	"example.com/auth-service/internal/logging"
	"example.com/auth-service/internal/utils/requestid"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		var (
			ip       = r.RemoteAddr
			method   = r.Method
			url_path = r.URL.String()
		)

		requestID, ok := requestid.GetRequestIDFromContext(r.Context())
		if !ok {
			requestID = "unknown"
		}

		rw := newResponseWriter(w)
		next.ServeHTTP(rw, r)

		elapsed := time.Since(start)

		logging.Logger.Info(
			"request received",
			logging.Group(
				"user",
				"ip", strings.Split(ip, ":")[0],
			),
			logging.Group(
				"request",
				"request_id", requestID,
				"status", rw.statusCode,
				"method", method,
				"url_path", url_path,
				"duration", elapsed.Milliseconds(),
				"human_readable_duration", elapsed.String(),
			),
		)

	})
}
