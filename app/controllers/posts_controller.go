package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pedroacl/todo-list/app/helpers"
	"github.com/pedroacl/todo-list/app/models"
)

func init() {
}

// GetPosts returns a list of posts
func GetPosts(w http.ResponseWriter, r *http.Request) {
	posts := []models.Post{}
	posts = append(posts, models.Post{Title: "Post 1"})
	posts = append(posts, models.Post{Title: "Post 2"})

	helpers.CreateJSONResponse(posts, http.StatusOK, w)
}

// GetPost returns a post based on its id
func GetPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params["id"])

	post := models.Post{Title: "Post " + params["id"]}

	helpers.CreateJSONResponse(post, http.StatusOK, w)
}

// UpdatePost shows an edit post form
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params["id"])

	msg := Message{"Post updated!"}
	helpers.CreateJSONResponse(msg, http.StatusOK, w)
}

// CreatePost creates a new post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		fmt.Println(err)
	}

	var post models.Post
	err = decoder.Decode(&post, r.PostForm)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("The post:", post)
}

// DeletePost deletes a post based on its id
func DeletePost(w http.ResponseWriter, r *http.Request) {
	// TODO
}
