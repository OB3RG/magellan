package author

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"

	"magellan/pkg/config"
	"magellan/pkg/middleware"
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

	handler := NewHandler()

	s.r.Get("/", middleware.IsAdmin(handler.handleIndex))

	return s.r
}
