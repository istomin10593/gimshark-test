package server

import (
	"context"
	"errors"
	"fmt"
	"gimshark-test/ui/internal/handler/pack_handler"
	"gimshark-test/ui/pkg/config"
	"net/http"
	"time"

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

	log.Info("Starting the UI server", zap.String("port", cfg.HTTP.Port))

	errCh := make(chan error, 1)

	go func() {
		defer close(errCh)

		<-ctx.Done()

		log.Info("server.Serve: context closed: shutting down")
		tx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		log.Info("server.Serve: shutting down")
		errCh <- srv.Shutdown(tx)
	}()

	// Run the server. This will block until the provided context is closed.
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Error("failed to serve", zap.Error(err))

		return fmt.Errorf("failed to serve: %w", err)
	}

	select {
	case err := <-errCh:
		log.Error("failed to shutdown", zap.Error(err))

		return fmt.Errorf("failed to shutdown: %w", err)
	default:
		return nil
	}
}

// routes returns the server routes.
func routes(log *zap.Logger, cfg *config.Config) http.Handler {
	// Create new instance of mux.
	r := mux.NewRouter()

	// Packs.
	packsRouter := r.PathPrefix("/packs").Subrouter()

	packHandler := pack_handler.New(pack_handler.NewConfig(cfg), log)
	packsRouter.HandleFunc("", packHandler.Index).Methods("GET")
	packsRouter.HandleFunc("/calculate", packHandler.Calculator).Methods("POST")

	return r
}
