package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pedroacl/todo-list/config"
)

func about(w http.ResponseWriter, r *http.Request) {
}

func main() {
	port := ":8080"

	r := config.GetRoutes()
	http.Handle("/", r)

	fmt.Println("Server started!")
	fmt.Println("Listening on port " + port)

	log.Fatal(http.ListenAndServe(port, nil))
}
