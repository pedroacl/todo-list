package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func init() {

}

// GetRoutes for the app
func GetRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/about", About).Methods("GET")
	r.HandleFunc("/getAuthToken", GetAuthToken).Methods("GET")

	// auth
	r.HandleFunc("/auth", GetAuthToken).Methods("GET")

	// users
	r.HandleFunc("/users", ValidateJWT(GetUsers)).Methods("GET")
	r.HandleFunc("/users/{id}", ValidateJWT(GetUser)).Methods("GET")
	r.HandleFunc("/users/{id}", ValidateJWT(UpdateUser)).Methods("PUT")
	r.HandleFunc("/users/{id}", ValidateJWT(DeleteUser)).Methods("DELETE")
	r.HandleFunc("/users/new", ValidateJWT(CreateUser)).Methods("POST")

	// posts
	r.HandleFunc("/posts", GetPosts).Methods("GET")
	r.HandleFunc("/posts", CreatePost).Methods("POST")
	r.HandleFunc("/posts/{id}", UpdatePost).Methods("PUT")

	//
	r.HandleFunc("/test", GetPosts).Methods("GET")

	// serve main index
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	return r
}
