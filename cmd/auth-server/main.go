package main

import (
	"log"
	"net/http"

	"example.com/auth-service/internal/api"
	"example.com/auth-service/internal/config"
)

func main() {
	cfg, err := config.Load("configs/local.yaml")
	if err != nil {
		log.Fatalf("fatal: config error: %v", err)
	}

	router := api.CreateRouter()

	log.Printf("Auth server starting on port %d...\n", cfg.Server.Port)
	log.Fatal(http.ListenAndServe(cfg.Addr(), router))
}
