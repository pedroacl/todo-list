package config

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pedroacl/todo-list/controllers"
)

func init() {

}

// GetRoutes for the app
func GetRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/about", controllers.About).Methods("GET")

	// users
	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	r.HandleFunc("/users/{userId}", controllers.GetUser).Methods("GET")
	r.HandleFunc("/users/{userId}", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{userId}", controllers.DeleteUser).Methods("DELETE")
	r.HandleFunc("/users/new", controllers.CreateUser).Methods("POST")

	// posts
	r.HandleFunc("/posts", controllers.GetPosts).Methods("GET")

	// serve main index
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	return r
}
