package main

import (
	"net/http"
	"encoding/json"
	"github.com/go-errors/errors"
	"github.com/gorilla/mux"
)

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

var userStore = []User{}

func main() {
	r := SetUserRoutes()
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	server.ListenAndServe()
}

func SetUserRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users", getUsers).Methods("GET")
	return r
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := json.Marshal(userStore)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "appliction/json")
	w.WriteHeader(http.StatusOK)
	w.Write(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User;
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = validate(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userStore = append(userStore, user)
	w.WriteHeader(http.StatusCreated)
}

func validate(user User) error {
	for _, u := range userStore {
		if u.Email == user.Email {
			return errors.New("The email already exists")
		}
	}
	return nil
}
