package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"gimshark-test/server/internal/handler/health_handler"
	"gimshark-test/server/pkg/config"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// Run starts the server.
func Run(ctx context.Context, log *zap.Logger, cfg *config.Config) error {
	srv := &http.Server{
		Handler:           routes(log, cfg),
		Addr:              cfg.HTTP.Port,
		ReadHeaderTimeout: cfg.HTTP.ReadHeaderTimeout,
	}

	log.Info("Starting the server", zap.String("port", cfg.HTTP.Port))

	errCh := make(chan error, 1)

	go func() {
		<-ctx.Done()

		log.Info("server.Serve: context closed: shutting down")
		tx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		log.Info("server.Serve: shutting down")
		errCh <- srv.Shutdown(tx)
	}()

	// Run the server. This will block until the provided context is closed.
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("failed to serve: %w", err)
	}

	select {
	case err := <-errCh:
		close(errCh)

		return fmt.Errorf("failed to shutdown: %w", err)
	default:
		return nil
	}
}

// routes returns the server routes.
func routes(log *zap.Logger, cfg *config.Config) http.Handler {
	// Create new instance of mux
	r := mux.NewRouter()

	// Health.
	healthHandler := health_handler.New()
	r.HandleFunc("/", healthHandler.Hello().ServeHTTP).Methods("GET")

	return r
}
