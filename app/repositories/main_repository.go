package repositories

import (
	"gopkg.in/mgo.v2"
)

// GetSession return a mongodb session
func GetSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost:27017/todo_list")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}

	return s
}
