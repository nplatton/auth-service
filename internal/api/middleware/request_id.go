package middleware

import (
	"net/http"

	"github.com/google/uuid"

	"example.com/auth-service/internal/utils/requestid"
)

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()

		ctx := requestid.SetRequestIDInContext(r.Context(), id)
		w.Header().Set("X-Request-ID", id)

		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
