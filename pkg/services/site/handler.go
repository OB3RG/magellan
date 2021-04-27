package site

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

// Handler ...
type Handler struct {
	db *sql.DB
}

// NewHandler ...
func NewHandler(db *sql.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) handleList(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Minute))
	defer cancel()

	sites, err := List(ctx, h.db)
	if err != nil {
		log.Print(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sites)
}

func (h *Handler) handleGetOne(w http.ResponseWriter, r *http.Request) {
	siteID := chi.URLParam(r, "ID")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Minute))
	defer cancel()

	site, err := GetOne(ctx, h.db, siteID)
	if err != nil {
		log.Print(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(site)
}
