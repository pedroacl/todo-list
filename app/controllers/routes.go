package controllers

import (
	"context"
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

var mgoSession *mgo.Session

// MgoSessionHandler middleware
func MgoSessionHandler(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		//ctx := newContextWithMgoSession(req.Context(), req)
		session, err := mgo.Dial("127.0.0.1:27017")

		if err != nil {
			panic(err)
		}

		ctx := context.WithValue(req.Context(), mgoSession, session)
		defer session.Close()

		next.ServeHTTP(rw, req.WithContext(ctx))
	})
}

// GetRouter returns the router with its respective routes
func GetRouter() *mux.Router {
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
	r.Handle("/labels", MgoSessionHandler(ValidateJWT(GetLabels))).Methods("GET")
	r.Handle("/labels/{id}", MgoSessionHandler(ValidateJWT(GetLabelDetails))).Methods("GET")

	// serve main index
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	http.Handle("/", r)

	return r
}
