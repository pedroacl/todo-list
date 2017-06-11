package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/pedroacl/todo-list/config"
	"github.com/pedroacl/todo-list/controllers"
)

func main() {
	config.LoadConfiguration("local")

	r := controllers.GetRoutes()
	http.Handle("/", r)

	fmt.Println("Server started!")
	fmt.Println("Listening on port " + config.MainConfig.Port)

	http.ListenAndServe(config.MainConfig.Port, handlers.LoggingHandler(os.Stdout, r))
}
