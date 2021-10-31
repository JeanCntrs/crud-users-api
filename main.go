package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const (
	GET    string = "GET"
	POST   string = "POST"
	PUT    string = "PUT"
	DELETE string = "DELETE"
)

func main() {
	fmt.Println("Starting server...")

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("GET Method")
	}).Methods(GET)

	router.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, found := mux.Vars(r)["id"]

		if found {
			message := "ID found: " + id

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(message)

			return
		}

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(nil)
	}).Methods(GET)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("POST Method")
	}).Methods(POST)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("PUT Method")
	}).Methods(PUT)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("DELETE Method")
	}).Methods(DELETE)

	server := http.Server{
		Addr:              ":8080",
		Handler:           router,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	server.ListenAndServe()
}
