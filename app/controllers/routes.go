package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

func init() {

}

// Route interface
type Route struct {
	Path   string `json:"path"`
	Method string `json:"method"`
}

func loadRoutes(file string) []Route {
	configFile, err := os.Open(file)

	defer configFile.Close()

	if err != nil {
		fmt.Println(err.Error())
	}

	jsonParser := json.NewDecoder(configFile)

	var routes []Route
	jsonParser.Decode(&routes)

	return routes
}

func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost:27017/todo-list")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}

	return s
}

// GetRoutes for the app
func GetRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/about", About).Methods("GET")

	// get authentication token
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

	// labels
	// labelsController := NewLabelsController(getSession())
	r.HandleFunc("/labels", GetLabels).Methods("GET")
	r.HandleFunc("/labels/{id}", GetLabelDetails).Methods("GET")

	// serve main index
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	http.Handle("/", r)

	return r
}
