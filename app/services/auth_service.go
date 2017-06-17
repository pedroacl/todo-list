package services

import (
	"encoding/json"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pedroacl/todo-list/config"
	"time"
)

// Claims struct
type Claims struct {
	Username string `json:"username"`
	// recommended having
	jwt.StandardClaims
}

// getSecretKey returns our secret key
func getSecretKey(token *jwt.Token) (interface{}, error) {
	// validate the algorithm used
	_, ok := token.Method.(*jwt.SigningMethodHMAC)

	if !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}

	return []byte(config.MainConfig.SecretKey), nil
}

// ValidateAuthToken parses and validates a JWT token
func ValidateAuthToken(tokenString string) bool {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application. The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, getSecretKey)

	if err != nil {
		fmt.Println("Error parsing token: ", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		fmt.Println("The claims: ", claims)
		fmt.Println("Token is not valid:", err)

		return false
	}

	return true
}

// CreateAuthToken create a JWT token
func CreateAuthToken() []byte {
	username := "myusername"

	// Expires the token and cookie in 1 hour
	expireToken := time.Now().Add(time.Hour * 1).Unix()

	// We'll manually assign the claims but in production you'd insert values from a database
	claims := Claims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "localhost:9000",
		},
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	signedToken, err := token.SignedString([]byte(config.MainConfig.SecretKey))

	if err != nil {
		fmt.Println(err)
	}

	data := map[string]string{
		"token": signedToken,
	}

	tokenJSON, err := json.Marshal(data)

	if err != nil {
		fmt.Println(err)
	}

	return tokenJSON
}
