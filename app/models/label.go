package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Label type
type Label struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt"`
}
