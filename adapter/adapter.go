package adapter

import "net/http"

// HandlerAdapter type represents the signature of
// all our HandlerAdapter functions the can be
// used to adapt our http handlers to specific
// extended functionality like logging, obtaining
// a db session etc.
type HandlerAdapter func(http.Handler) http.Handler

// Adapt function is used to add the functionality
// of the n number of HandlerAdapter to the provided
// Handler
func AdaptHandler(h http.Handler, handlerAdapters ...HandlerAdapter) http.Handler {
	for _, handlerAdapters := range handlerAdapters {
		h = handlerAdapters(h)
	}
	return h
}
