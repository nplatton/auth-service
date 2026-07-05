package handlers

import (
	"encoding/json"
	"net/http"

	"example.com/auth-service/internal/logging"
)

type HealthResponse struct {
	Status      string `json:"status"`
	Uptime      string `json:"uptime"`
	StartedAt   string `json:"started_at"`
	Version     string `json:"version"`
	Environment string `json:"environment"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := HealthResponse{
		Status: "healthy",
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		logging.Logger.Error(
			"failed to write response",
			"error", err,
			logging.Group(
				"request",
				"method", r.Method,
				"path", r.URL.Path,
			),
		)
	}
}
