package main

import (
	"testing"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strings"
)

func TestGetUsers(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers).Methods("GET")
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Errorf("Http Status expected: 200, got :%d", w.Code)
	}
}

func TestCreateUser(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/users", createUser).Methods("POST")
	userJson := `{"firstname":"shiju","lastname":"lucas","email":"lucas@ebet.com"}`
	req, err := http.NewRequest("POST", "/users", strings.NewReader(userJson))
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != 201 {
		t.Errorf("Http Status expected: 201, got %d", w.Code)
	}
}
