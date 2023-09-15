package pack_handler

import (
	"go.uber.org/zap"
)

// Handler defines a HTTP handler.
type Handler struct {
	cfg *Config
	log *zap.Logger
}

// New creates a new HTTP handler.
func New(
	cfg *Config,
	log *zap.Logger,
) *Handler {
	return &Handler{
		cfg: cfg,
		log: log,
	}
}
