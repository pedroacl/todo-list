package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pedroacl/todo-list/app/helpers"
	"github.com/pedroacl/todo-list/app/services"
	"net/http"
)

// GetLabels returns a list of labels
func GetLabels(w http.ResponseWriter, r *http.Request) {
	labels := services.GetLabels(getMgoSessionFromContext(r))
	fmt.Println(labels)
	helpers.CreateJSONResponse(labels, http.StatusOK, w)
}

// GetLabelDetails returns a label based on its ID
func GetLabelDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	labelID := params["id"]
	fmt.Println(labelID)

	label := services.GetLabel(getMgoSessionFromContext(r), labelID)
	helpers.CreateJSONResponse(label, http.StatusOK, w)
}
