package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func init() {
}

// Index function
func Index(w http.ResponseWriter, r *http.Request) {
	/*
		users := models.Users{
			models.User{Name: "Pedro"},
			models.User{Name: "Luis"},
		}

		json.NewEncoder(w).Encode(users)
	*/
	fmt.Println(r.URL.Path[1:])
	http.ServeFile(w, r, "./static/index.html")
	// http.FileServer(http.Dir("./static/index.html"))
}

// Extra function
func Extra(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome, %s!", r.URL.Path[1:])
}

// About lalala
func About(w http.ResponseWriter, r *http.Request) {
	m := Message{"Welcome"}
	b, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}

	w.Write(b)
}
