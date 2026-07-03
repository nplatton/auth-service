package api

import "net/http"

func CreateRouter() http.Handler {
	mux := http.NewServeMux()

	RegisterRoutes(mux)

	return mux
}
