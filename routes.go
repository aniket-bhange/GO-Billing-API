package main

import (
	"billing-gorilla/modules/users"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}

var routes = Routes{
	Route{
		"Get",
		"GET",
		"/users",
		users.Get,
	},
	Route{
		"Create",
		"POST",
		"/users/create",
		users.Create,
	},
	Route{
		"Update",
		"PUT",
		"/users/update/{id}",
		users.Update,
	},
}