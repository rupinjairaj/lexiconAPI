package database

import (
	"lexiconAPI/adapter"
	"net/http"

	"github.com/gorilla/context"
	"gopkg.in/mgo.v2"
)

// WithDataBase is an Adapter function
// that extends the functionality of the
// provided handler by adding a mongodb
// session object to it.
func WithDataBase(db *mgo.Session) adapter.HandlerAdapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			dbsession := db.Copy()
			defer dbsession.Close()

			context.Set(r, "database", dbsession)
			h.ServeHTTP(w, r)
		})
	}
}
