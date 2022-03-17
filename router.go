package main

import (
	"github.com/gorilla/mux"
	"ApiProject/controllers"
)

func InitializeRouter() *mux.Router {
	// StrictSlash is true => redirect /oeuvres/ to /oeuvres
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET").Path("/oeuvres").Name("Index").HandlerFunc(controllers.OeuvresIndex)
	router.Methods("POST").Path("/cars").Name("Create").HandlerFunc(controllers.OeuvresCreate)
	router.Methods("GET").Path("/cars/{id}").Name("Show").HandlerFunc(controllers.OeuvresShow)
	router.Methods("PUT").Path("/cars/{id}").Name("Update").HandlerFunc(controllers.OeuvresUpdate)
	router.Methods("DELETE").Path("/cars/{id}").Name("DELETE").HandlerFunc(controllers.OeuvresDelete)
	return router
}