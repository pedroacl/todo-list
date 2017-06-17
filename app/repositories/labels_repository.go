package repositories

import (
	"github.com/pedroacl/todo-list/app/models"
	"gopkg.in/mgo.v2"
)

func labelsCollection(mgoSession *mgo.Session) *mgo.Collection {
	return mgoSession.DB("todo_list").C("labels")
}

// GetLabels returns a list of labels
func GetLabels(mgoSession *mgo.Session) []models.Label {
	var labels []models.Label
	labelsCollection(mgoSession).Find(nil).All(&labels)

	// labels = append(labels, models.Label{Name: "Label 1"})
	// labels = append(labels, models.Label{Name: "Label 2"})

	return labels
}

// GetLabelByID returns a single label
func GetLabelByID(mgoSession *mgo.Session, labelID string) models.Label {
	var label models.Label
	labelsCollection(mgoSession).Find(labelID).One(&label)

	return label
}
