package models

import (
	"gopkg.in/mgo.v2/bson"
)

// Label type
type Label struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	Name        string        `json:"id"`
	Description string        `json:"description"`
}
