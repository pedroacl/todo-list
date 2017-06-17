package services

import (
	"github.com/pedroacl/todo-list/app/models"
	"github.com/pedroacl/todo-list/app/repositories"
)

func init() {

}

// GetLabels returns a list of posts
func GetLabels() []models.Label {
	labels := repositories.GetLabels()

	return labels
}
