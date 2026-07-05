package api

import (
	"encoding/json"
	"net/http"

	"example.com/auth-service/internal/logging"
	"example.com/auth-service/internal/requestid"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func WriteError(w http.ResponseWriter, r *http.Request, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	requestID, ok := requestid.GetRequestIDFromContext(r.Context())
	if !ok {
		requestID = "unknown"
	}

	if err := json.NewEncoder(w).Encode(ErrorResponse{Error: message}); err != nil {
		logging.Logger.Error(
			"failed to write error response",
			"request_id", requestID,
			"error", err,
			logging.Group(
				"request",
				"method", r.Method,
				"path", r.URL.Path,
			),
		)
	}
}
