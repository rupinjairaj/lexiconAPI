package word

import (
	"log"
	"net/http"
)

type Controller struct {
	Repository Repository
}

func (c *Controller) Word(w http.ResponseWriter, r *http.Request) {
	wordData := c.Repository.GetWordData(w, r)
	log.Println(string(wordData))
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(wordData)
}
