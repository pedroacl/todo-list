package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/pedroacl/todo-list/config"
)

func main() {
	port := ":8080"

	r := config.GetRoutes()
	http.Handle("/", r)

	fmt.Println("Server started!")
	fmt.Println("Listening on port " + port)

	http.ListenAndServe(port, handlers.LoggingHandler(os.Stdout, r))
}
