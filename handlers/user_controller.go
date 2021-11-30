package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/JeanCntrs/api-intermediate-level/entities"
	"github.com/JeanCntrs/api-intermediate-level/services"
	"github.com/gorilla/mux"
)

const (
	GET    string = "GET"
	POST   string = "POST"
	PUT    string = "PUT"
	DELETE string = "DELETE"
)

type UserController interface {
	HandleGeneral(w http.ResponseWriter, r *http.Request)
	HandleOne(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	service services.UserService
}

func NewUserController() UserController {
	return &userController{service: services.NewUserService()}
}

func (uc *userController) HandleGeneral(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == GET {
		users := uc.service.FindAll()
		json.NewEncoder(w).Encode(users)
	}

	if r.Method == POST {
		user := entities.User{}
		json.NewDecoder(r.Body).Decode(&user)

		uc.service.Create(user)
		json.NewEncoder(w).Encode(user)
	}

	if r.Method == DELETE {
		json.NewEncoder(w).Encode("DELETE Method")
	}
}

func (uc *userController) HandleOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == GET {
		id, found := mux.Vars(r)["id"]

		if found {
			user, err := uc.service.FindOne(id)

			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(err.Error())

				return
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(user)

			return
		}

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(nil)
	}

	if r.Method == PUT {
		id, found := mux.Vars(r)["id"]

		if found {
			user := entities.User{}
			json.NewDecoder(r.Body).Decode(&user)

			userUpdated := uc.service.Update(id, user)

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(userUpdated)

			return
		}

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("user not updated")
	}
}
