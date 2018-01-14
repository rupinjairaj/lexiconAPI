package main

import (
	"lexiconAPI/adapter"
	"lexiconAPI/database"
	"lexiconAPI/word"
	"log"
	"net/http"

	"fmt"

	"github.com/gorilla/handlers"
	"gopkg.in/mgo.v2"
)

var controller = &word.Controller{Repository: word.Repository{}}

func handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		controller.Word(w, r)
	default:
		http.Error(w, "Not supported", http.StatusMethodNotAllowed)
	}
}

func main() {
	db, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD"})

	h := adapter.AdaptHandler(http.HandlerFunc(handle), database.WithDatabase(db))

	server := &http.Server{
		Addr:    ":8808",
		Handler: handlers.CORS(allowedMethods, allowedOrigins)(h),
	}
	fmt.Println("Listening...")
	serverErr := server.ListenAndServe()
	if serverErr != nil {
		log.Fatalln(serverErr)
	}
}
