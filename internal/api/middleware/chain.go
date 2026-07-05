package middleware

import "net/http"

type Middleware func(next http.Handler) http.Handler

func Chain(middleware ...Middleware) func(http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		for i := len(middleware) - 1; i >= 0; i-- {
			handler = middleware[i](handler)
		}

		return handler
	}
}
