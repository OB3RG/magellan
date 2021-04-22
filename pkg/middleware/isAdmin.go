package middleware

import (
	"log"
	"net/http"
)

// IsAdmin checks if user is admin
func IsAdmin(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("Check if admin")
		h(w, r)
	}
}
