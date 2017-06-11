package config

import (
	"fmt"
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
	r.HandleFunc("/getAuthToken", controllers.GetAuthToken).Methods("GET")

	// auth
	r.HandleFunc("/auth", controllers.GetAuthToken).Methods("GET")

	// users
	r.HandleFunc("/users", validateJWT(controllers.GetUsers)).Methods("GET")
	r.HandleFunc("/users/{userId}", validateJWT(controllers.GetUser)).Methods("GET")
	r.HandleFunc("/users/{userId}", validateJWT(controllers.UpdateUser)).Methods("PUT")
	r.HandleFunc("/users/{userId}", validateJWT(controllers.DeleteUser)).Methods("DELETE")
	r.HandleFunc("/users/new", validateJWT(controllers.CreateUser)).Methods("POST")

	// posts
	r.HandleFunc("/posts", controllers.GetPosts).Methods("GET")

	//
	r.HandleFunc("/test", controllers.GetPosts).Methods("GET")

	// serve main index
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	return r
}

// Middleware to protect /profile and /logout
func validateJWT(protectedPage http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwtToken := r.Header.Get("Auth-Token")
		fmt.Println("The token: ", jwtToken)

		// validate token
		// isTokenValid := controllers.ValidateAuthToken(jwtToken)
		isTokenValid := true

		//Validate the token and if it passes call the protected handler below.
		if isTokenValid {
			protectedPage(w, r)
		}
	})
}
