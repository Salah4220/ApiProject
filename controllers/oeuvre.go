package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"ApiProject/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func OeuvresIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(models.AllOeuvre())
}

func OeuvresCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var oeuvre models.Oeuvre

	err = json.Unmarshal(body, &oeuvre)

	if err != nil {
		log.Fatal(err)
	}

	models.NewOeuvre(&oeuvre)

	json.NewEncoder(w).Encode(oeuvre)
}

func OeuvresShow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Fatal(err)
	}

	oeuvre := models.FindOeuvreById(id)

	json.NewEncoder(w).Encode(oeuvre)
}
func OeuvresUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	oeuvre := models.FindOeuvreById(id)

	err = json.Unmarshal(body, &oeuvre)

	models.UpdateOeuvre(oeuvre)

	json.NewEncoder(w).Encode(oeuvre)
}
func OeuvresDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)

	// strconv.Atoi is shorthand for ParseInt
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Fatal(err)
	}

	err = models.DeleteOeuvreById(id)
}