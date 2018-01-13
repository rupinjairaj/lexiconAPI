package logger

import (
	"lexiconAPI/adapter"
	"log"
	"net/http"
	"time"
)

// Logger is a wrapper around the handlers used
// throughout the api to insert logs for every
// request received by the api
func Logger(logger *log.Logger) adapter.HandlerAdapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			h.ServeHTTP(w, r)
			logger.Printf("%s\t%s\t%s", r.Method, r.RequestURI, time.Since(start))
		})
	}
}
