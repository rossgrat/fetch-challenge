package logger

import (
	"log"
	"net/http"
)

// Basic logging with the remote address, method, path, and an optional message
// If the request is not defined, just log the message
func LogInfo(r *http.Request, msg string) {
	if r != nil {
		log.Printf("%s %s %s %s", r.RemoteAddr, r.Method, r.URL.Path, msg)
	} else {
		log.Printf("%s", msg)
	}
}
