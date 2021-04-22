package tag

import "net/http"

// Handler ...
type Handler struct{}

// NewHandler ...
func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Tag hi"))
}
