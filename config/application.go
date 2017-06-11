package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config struct
type Config struct {
	Port      string `json:"port"`
	SecretKey string `json:"secretKey"`
}

// MainConfig the main configuration struct
var MainConfig Config

// LoadConfiguration load
func LoadConfiguration(env string) {
	var config Config
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
	jsonParser.Decode(&config)

	MainConfig = config
}
