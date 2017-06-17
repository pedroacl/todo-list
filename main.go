package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/pedroacl/todo-list/app/controllers"
	"github.com/pedroacl/todo-list/config"
)

func main() {
	config.LoadConfiguration("local")
	router := controllers.GetRoutes()

	fmt.Println("Server started!")
	fmt.Println("Listening on port " + config.MainConfig.Port)

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	log.Fatal(http.ListenAndServe(config.MainConfig.Port,
		handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
