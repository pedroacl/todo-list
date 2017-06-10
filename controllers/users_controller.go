package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	models "github.com/pedroacl/todo-list/models"
)

func init() {
	// var decoder = schema.NewDecoder()
}

// Message structure
type Message struct {
	Text string
}

// GetUsers returns a list of users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := models.Users{
		models.User{Name: "Pedro"},
		models.User{Name: "Luis"},
	}

	json.NewEncoder(w).Encode(users)
}

// GetUser returns a user based on his id
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Println(vars)
	fmt.Println(r.Method)

	m := Message{"Welcome user!"}
	b, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

// UpdateUser updates a given user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	m := Message{"Welcome user!"}
	b, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

// CreateUser creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {

}

// DeleteUser deletes a user based on his id
func DeleteUser(w http.ResponseWriter, r *http.Request) {

}
