package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"example.com/auth-service/internal/api"
	m "example.com/auth-service/internal/api/middleware"
	"example.com/auth-service/internal/config"
)

func main() {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "configs/local.yaml"
	}

	cfg, err := config.Load(path)
	if err != nil {
		log.Fatalf("fatal: config error: %v", err)
	}

	router := api.CreateRouter()

	server := &http.Server{
		Addr:    cfg.Addr(),
		Handler: m.Chain(m.Recover, m.RequestID, m.Logger)(router),
	}

	serverError := make(chan error, 1)

	log.Printf("Auth server starting on %s...\n", server.Addr)

	go func() {
		serverError <- server.ListenAndServe()
	}()

	sigChan := waitForShutdownSignal()

	select {
	case err := <-serverError:
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v", err)
		}

	case <-sigChan:
		log.Println("Shutdown signal received.")
		shutdownCtx, shutdownRelease := context.WithTimeout(
			context.Background(),
			10*time.Second,
		)
		defer shutdownRelease()

		shutdownServer(server, shutdownCtx)
	}

}

func waitForShutdownSignal() <-chan os.Signal {
	sigChan := make(chan os.Signal, 1)

	signal.Notify(
		sigChan,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	return sigChan
}

func shutdownServer(server *http.Server, ctx context.Context) {
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("HTTP shutdown error: %v", err)
	}
	log.Println("Server shut down gracefully.")
}
