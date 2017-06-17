package config

import (
	"encoding/json"
	"fmt"
	//"github.com/urfave/negroni"
	"os"
)

// Config struct
type Config struct {
	Port      string `json:"port"`
	SecretKey string `json:"secretKey"`
}

// MainConfig the main configuration struct
var MainConfig Config

func getMainConfig(env string) Config {
	var file string

	if env == "local" {
		file = "config/env.local.json"
	} else {
		file = "config/env.local.json"
	}

	configFile, err := os.Open(file)
	defer configFile.Close()

	if err != nil {
		fmt.Println(err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	var config Config
	jsonParser.Decode(&config)

	return config
}

/*
// MongoMiddleware middleware for handling a mgo session
func MongoMiddleware() negroni.HandlerFunc {
	// database := os.Getenv("DB_NAME")
	database := "todo_list"
	session, err := mgo.Dial("127.0.0.1:27017")

	if err != nil {
		panic(err)
	}

	return negroni.HandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		reqSession := session.Clone()
		defer reqSession.Close()

		ctx := r.Context().WithValue(r.Context(), mgoSession, reqSession)
		next(rw, r.WithContext(ctx))
	})
}
*/

// LoadConfiguration load
func LoadConfiguration(env string) {
	MainConfig = getMainConfig("local")
}
