package controllers

import (
	"fmt"
	"net/http"

	"github.com/pedroacl/todo-list/app/helpers"
	"github.com/pedroacl/todo-list/app/services"
)

// GetJWTHandler return a JWT token
func GetJWTHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Validate the user with a DB before creating token
	token := services.CreateAuthToken()

	helpers.CreateJSONResponse(token, http.StatusOK, w)
}

// ValidateJWTHandler middleware to protect /profile and /logout
func ValidateJWTHandler(protectedPage http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwtToken := r.Header.Get("Auth-Token")
		fmt.Println("The token: ", jwtToken)

		// validate the token and if it passes call the protected handler below.
		// isTokenValid := services.ValidateAuthToken(jwtToken)
		isTokenValid := true

		if isTokenValid {
			protectedPage(w, r)
		}
	})
}
