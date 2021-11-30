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

	userController := handlers.NewUserController()
	router := mux.NewRouter()

	router.HandleFunc("/", userController.HandleGeneral)
	router.HandleFunc("/{id}", userController.HandleOne)

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
