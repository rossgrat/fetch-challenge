package mw

import (
	"net/http"
	"strconv"

	"github.com/rossgrat/fetch-challenge/src/logger"
)

// Basic middleware for logging HTTP transactions as they come int
func LogRequest(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ww := &WrappedResponseWriter{
			W:          w,
			StatusCode: 0,
		}
		f(ww, r)
		ww.Done()
		logger.LogInfo(r, strconv.Itoa(ww.StatusCode))
	}
}

// The http.ResponseWriter in Golang has the annoying quality
// of being unable to access the status code once set. This is
// just a simple wrapper that saves the status code when set
// to the underlying response writer
type WrappedResponseWriter struct {
	W          http.ResponseWriter
	StatusCode int
}

func (ww *WrappedResponseWriter) Header() http.Header {
	return ww.W.Header()
}
func (ww *WrappedResponseWriter) Write(b []byte) (int, error) {
	return ww.W.Write(b)
}
func (ww *WrappedResponseWriter) WriteHeader(statusCode int) {
	ww.StatusCode = statusCode
	ww.W.WriteHeader(statusCode)
}
func (ww *WrappedResponseWriter) Done() {
	if ww.StatusCode == 0 {
		ww.StatusCode = http.StatusOK
	}
}
