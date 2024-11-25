package mw

import (
	"log"
	"net/http"
)

// Basic middleware for logging HTTP transactions as they come int
func LogRequest(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Host, r.Method, r.URL.Path)
		f(w, r)
	}
}
