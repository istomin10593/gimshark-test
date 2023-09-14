package pack_handler

import (
	"net/http"
	"text/template"

	"go.uber.org/zap"
)

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		h.log.Error("index.Template", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		h.log.Error("index.Template", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
