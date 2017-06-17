package controllers

import (
	"fmt"
	"github.com/pedroacl/todo-list/app/helpers"
	"github.com/pedroacl/todo-list/app/services"
	//"gopkg.in/mgo.v2"
	"net/http"
)

func init() {

}

// GetLabels returns a list of labels
func GetLabels(w http.ResponseWriter, r *http.Request) {
	labels := services.GetLabels(getMgoSessionFromContext(r))
	fmt.Println(labels)
	helpers.CreateJSONResponse(labels, http.StatusOK, w)
}

// GetLabelDetails returns a label based on its ID
func GetLabelDetails(w http.ResponseWriter, r *http.Request) {

}
