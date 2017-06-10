package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pedroacl/todo-list/helpers"
	"github.com/pedroacl/todo-list/models"
)

func init() {
}

// GetPosts returns a list of posts
func GetPosts(w http.ResponseWriter, r *http.Request) {
	users := models.Users{
		models.User{Name: "Pedro"},
		models.User{Name: "Luis"},
	}

	// json.NewEncoder(w).Encode(users)
	usersJSON, err := json.Marshal(users)

	if err != nil {
		msg := Message{"Error getting posts"}
		data, err := json.Marshal(msg)

		if err != nil {
			// TODO
		}

		helpers.CreateJSONResponse(data, http.StatusInternalServerError, w)
		return
	}

	helpers.CreateJSONResponse(usersJSON, http.StatusOK, w)
}

// GetPost returns a post based on its id
func GetPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Println(vars)
	fmt.Println(r.Method)

	m := Message{"Welcome user!"}
	b, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}

	helpers.CreateJSONResponse(b, http.StatusOK, w)
}

// UpdatePost shows an edit post form
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	m := Message{"Welcome user!"}
	b, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

// CreatePost creates a new post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		// TODO handle error
	}

	post := new(models.Post)
	err = decoder.Decode(post, r.PostForm)

	if err != nil {
		// TODO handle error
	}

	fmt.Println(post)
}

// DeletePost deletes a post based on its id
func DeletePost(w http.ResponseWriter, r *http.Request) {

}
