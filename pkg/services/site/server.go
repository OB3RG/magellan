package site

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"

	"magellan/pkg/config"
)

// Server ...
type Server struct {
	cfg *config.Config
	db  *sql.DB
	r   *chi.Mux
}

// NewServer ...
func NewServer(cfg *config.Config, db *sql.DB) *Server {
	return &Server{cfg: cfg, db: db}
}

// SetupRoutes ...
func (s *Server) SetupRoutes() http.Handler {
	s.r = chi.NewRouter()

	handler := NewHandler(s.db)

	s.r.Get("/", handler.handleList)
	s.r.Get("/{ID}", handler.handleGetOne)

	return s.r
}
