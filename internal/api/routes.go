package api

import (
	"net/http"

	"example.com/auth-service/internal/api/handlers"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", handlers.HealthHandler)
}
