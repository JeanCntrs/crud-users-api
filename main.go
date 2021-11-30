package main

import (
	"github.com/gorilla/mux"

	"github.com/JeanCntrs/api-intermediate-level/handlers"

	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Starting server...")

	handlerGeneral := handlers.NewUserController()
	router := mux.NewRouter()
	router.HandleFunc("/", handlerGeneral.HandleGeneral)
	router.HandleFunc("/{id}", handlerGeneral.HandleOne)

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
