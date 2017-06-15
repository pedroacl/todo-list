package controllers

import (
	"github.com/pedroacl/todo-list/app/services"
	"net/http"
)

func init() {

}

// GetLabels returns a list of labels
func GetLabels(w http.ResponseWriter, r *http.Request) {
	services.GetLabels()
}

// GetLabelDetails returns a label based on its ID
func GetLabelDetails(w http.ResponseWriter, r *http.Request) {

}
