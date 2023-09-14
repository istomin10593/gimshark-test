package pack_handler

import (
	"encoding/json"
	"errors"
	"net/http"

	pack_usecase "gimshark-test/server/internal/usecase/pack"

	"go.uber.org/zap"
)

// Requst validation errors.
var ErrZeroItems = errors.New("items should be greater than 0")

// Handler defines a HTTP handler.
type Handler struct {
	uCase *pack_usecase.Usecase
	log   *zap.Logger
}

// New creates a new HTTP handler.
func New(
	log *zap.Logger,
	uCase *pack_usecase.Usecase,
) *Handler {
	return &Handler{
		uCase: uCase,
		log:   log,
	}
}

// In is dto for http req.
type GetPacksNumberIn struct {
	Items uint64 `json:"items"`
}

// validates request.
func (h Handler) validateReq(in *GetPacksNumberIn) error {
	if in.Items == 0 {
		return ErrZeroItems
	}

	return nil
}

// Create responsible for saving new order.
func (h Handler) GetPacksNumber() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		in := &GetPacksNumberIn{}
		// parse req body to dto
		err := json.NewDecoder(r.Body).Decode(&in)
		if err != nil {
			h.log.Error("can't parse req", zap.Error(err))
			http.Error(w, "Bad request: invalid data", http.StatusBadRequest)

			return
		}

		// check that request valid
		err = h.validateReq(in)
		if err != nil {
			h.log.Error("bad req", zap.Error(err))
			http.Error(w, "Bad request: "+err.Error(), http.StatusBadRequest)

			return
		}

		packsSize := h.uCase.GetPacksNumber(in.Items)

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(packsSize); err != nil {
			h.log.Error("packsNum.Encode", zap.Error(err))
			http.Error(w, "Internal server error", http.StatusInternalServerError)

			return
		}
	}

	return http.HandlerFunc(fn)
}
