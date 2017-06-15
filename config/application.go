package config

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"os"
)

// Config struct
type Config struct {
	Port      string `json:"port"`
	SecretKey string `json:"secretKey"`
}

// MainConfig the main configuration struct
var MainConfig Config

// MongoDBSession database session
var MongoDBSession *mgo.Session

func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost:27017/todo-list")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}

	return s
}

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

// LoadConfiguration load
func LoadConfiguration(env string) {
	MainConfig = getMainConfig("local")
	MongoDBSession = getSession()
}
