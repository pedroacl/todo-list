package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/pedroacl/todo-list/app/controllers"
	"github.com/pedroacl/todo-list/config"
)

func main() {
	config.LoadConfiguration("local")

	router := controllers.GetRouter()

	// Gorilla CORS
	mainHandlers := handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"}),
		handlers.AllowedHeaders([]string{"X-Requested-With"}))(router))

	// Server main config
	s := &http.Server{
		Addr:           config.MainConfig.Port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Start server
	fmt.Println("Server started!")
	fmt.Println("Listening on port " + config.MainConfig.Port)
	log.Fatal(s.ListenAndServe(), mainHandlers)
}
