package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"time"
	"fmt"
)

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}

func New() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", Hello).Methods("GET")
	router.HandleFunc("/api", Hello).Methods("GET")
	return router
}

func main() {
	srv := &http.Server{
			Handler: New(),
			Addr:    "0.0.0.0:8000",
			// Good practice: enforce timeouts for servers you create!
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}