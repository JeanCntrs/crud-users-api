package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	GET    string = "GET"
	POST   string = "POST"
	PUT    string = "PUT"
	DELETE string = "DELETE"
)

type InterfaceGeneral interface {
	HandleGeneral(w http.ResponseWriter, r *http.Request)
	HandleOne(w http.ResponseWriter, r *http.Request)
}

type HandlerGeneral struct{}

func NewHandlerGeneral() InterfaceGeneral {
	return HandlerGeneral{}
}

func (hg HandlerGeneral) HandleGeneral(w http.ResponseWriter, r *http.Request) {
	if r.Method == GET {
		json.NewEncoder(w).Encode("GET Method")
	}

	if r.Method == POST {
		json.NewEncoder(w).Encode("POST Method")
	}

	if r.Method == PUT {
		json.NewEncoder(w).Encode("PUT Method")
	}

	if r.Method == DELETE {
		json.NewEncoder(w).Encode("DELETE Method")
	}
}

func (hg HandlerGeneral) HandleOne(w http.ResponseWriter, r *http.Request) {
	if r.Method == GET {
		id, found := mux.Vars(r)["id"]

		if found {
			message := "ID found: " + id

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(message)

			return
		}

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(nil)
	}
}
