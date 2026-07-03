// Application entry point responsible for loading configuration, wiring dependencies and starting the HTTP server.

package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/auth-service/internal/api"
)

func main() {
	router := api.CreateRouter()

	// Start the server on port 8080
	fmt.Println("Auth server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
