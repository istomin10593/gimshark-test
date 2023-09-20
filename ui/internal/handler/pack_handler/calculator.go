package pack_handler

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"text/template"
	"time"

	"go.uber.org/zap"
)

const (
	maxRetries = 3
)

// GetPacksNumberIn - input for get packs number.
type GetPacksNumberOut struct {
	Items int `json:"items"`
}

// Calculator - handler for calculating the number and size of packs.
// Views result page.
func (h *Handler) Calculator(w http.ResponseWriter, r *http.Request) {
	quantity, err := strconv.Atoi(r.FormValue("quantity"))
	if err != nil {
		h.log.Error("invalid quantity", zap.Error(err))
		http.Error(w, "Invalid quantity", http.StatusBadRequest)

		return
	}
	// Define the request body.
	reqBody := GetPacksNumberOut{
		Items: quantity,
	}

	// Marshal the request body to JSON.
	reqJSON, err := json.Marshal(reqBody)
	if err != nil {
		h.log.Error("error marshaling request body", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)

		return
	}

	// Create the request to the backend API.
	u := &url.URL{
		Scheme: "http",
		Host:   h.cfg.Host + h.cfg.Port,
		Path:   h.cfg.Endpoint,
	}

	// Make the request to the backend API.
	for attempt := 0; attempt < maxRetries; attempt++ {
		// Create a context with a timeout.
		ctx, cancel := context.WithTimeout(context.Background(), h.cfg.Timeout)
		defer cancel()

		// Create the request to the backend API.
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, u.String(), bytes.NewBuffer(reqJSON))
		if err != nil {
			h.log.Error("calculator.Request", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			h.log.Error("calculator.Response", zap.Error(err))
			if attempt == maxRetries-1 {
				http.Error(w, err.Error(), http.StatusInternalServerError)

				return
			}

			// Calculate exponential backoff time
			exponentialBackoffTime := time.Duration(1<<attempt) * time.Second

			// Generate a random jitter value
			randomJitter := time.Duration(rand.Intn(int(h.cfg.MaxJitterWait-h.cfg.MinJitterWait))) + h.cfg.MinJitterWait

			// Apply exponential backoff with jitter
			waitTime := exponentialBackoffTime + randomJitter

			time.Sleep(waitTime)

			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				h.log.Error("calculator.Response", zap.Error(err))
				http.Error(w, "Error retrieving data from backend", http.StatusInternalServerError)

				return
			}

			http.Error(w, string(body), resp.StatusCode)

			return
		}

		// Parse the response JSON.
		packs := make(map[string]int)
		if err := json.NewDecoder(resp.Body).Decode(&packs); err != nil {
			h.log.Error("calculator.Encode", zap.Error(err))
			http.Error(w, "Internal server error", http.StatusInternalServerError)

			return
		}

		tmpl, err := template.ParseFiles("templates/result.html")
		if err != nil {
			h.log.Error("calculator.Template", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		if err := tmpl.Execute(w, packs); err != nil {
			h.log.Error("calculator.Template", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}
}
