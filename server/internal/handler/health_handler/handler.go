package health_handler

import (
	"fmt"
	"net/http"
)

// Handler defines a HTTP handler.
type Handler struct {
}

// New creates a new HTTP handler.
func New() *Handler {
	return &Handler{}
}

// hello checks the health of the service.
func (h *Handler) Hello() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		if _, err := w.Write([]byte("Hello, World!")); err != nil {
			fmt.Printf("err: %v", err)
		}
	}

	return http.HandlerFunc(fn)
}
