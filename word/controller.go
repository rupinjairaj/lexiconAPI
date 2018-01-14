package word

import (
	"net/http"
)

// Controller :
type Controller struct {
	Repository Repository
}

// Word :
func (c *Controller) Word(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	c.Repository.GetWordData(w, r)
}
