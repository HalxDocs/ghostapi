package api

import (
	"encoding/json"
	"net/http"

	"github.com/halxdocs/ghostapi/internal/engine"
)

type Handler struct {
	engine *engine.Engine
}

func NewHandler() *Handler {
	return &Handler{
		engine: engine.NewEngine(),
	}
}

type Request struct {
	URL string `json:"url"`
}

func (h *Handler) Scrape(w http.ResponseWriter, r *http.Request) {
	var req Request

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.URL == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	data, err := h.engine.Process(req.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}