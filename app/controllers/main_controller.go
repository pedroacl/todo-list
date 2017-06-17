package controllers

import (
	"github.com/gorilla/schema"
	"gopkg.in/mgo.v2"
	"net/http"
)

var decoder = schema.NewDecoder()

func getMgoSessionFromContext(r *http.Request) *mgo.Session {
	return r.Context().Value(mgoSession).(*mgo.Session)
}
