package services

import (
	"github.com/pedroacl/todo-list/app/models"
	"github.com/pedroacl/todo-list/app/repositories"
	"gopkg.in/mgo.v2"
)

// GetLabels returns a list of posts
func GetLabels(mgoSession *mgo.Session) []models.Label {
	labels := repositories.GetLabels(mgoSession)

	return labels
}

// GetLabel returns a label
func GetLabel(mgoSession *mgo.Session, labelID string) models.Label {
	label := repositories.GetLabelByID(mgoSession, labelID)

	return label
}
