package repositories

import (
	"github.com/pedroacl/todo-list/app/models"
)

func init() {

}

// GetLabels returns a list of labels
func GetLabels() []models.Label {
	var labels []models.Label

	session := GetSession()
	defer session.Close()

	session.DB("todo_list").C("labels").Find(nil).All(&labels)

	return labels
}
