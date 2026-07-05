package middleware

import (
	"net/http"
	"runtime/debug"
	"strings"

	"example.com/auth-service/internal/api"
	"example.com/auth-service/internal/logging"
)

func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if p := recover(); p != nil {
				logPanic(r, p)
				api.WriteError(w, r, http.StatusInternalServerError, "Internal Server Error")
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func logPanic(r *http.Request, err any) {
	logging.Logger.Error(
		"panic recovered",
		logging.Group(
			"panic",
			"error", err,
			"stack_trace", string(debug.Stack()),
		),
		logging.Group(
			"user",
			"ip", strings.Split(r.RemoteAddr, ":")[0],
		),
		logging.Group(
			"request",
			"method", r.Method,
			"path", r.URL.Path,
		),
	)
}
