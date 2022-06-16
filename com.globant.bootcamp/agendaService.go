package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func getAgendas(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Agenda Data:\n")
	for _, agenda := range getDB() {
		json.NewEncoder(w).Encode(agenda)
	}
	w.WriteHeader(http.StatusAccepted)
}

func saveAgenda(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var agenda Agenda
	json.Unmarshal(requestBody, &agenda)

	updateData(append(getDB(), agenda))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(agenda)
}

func getById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])
	agenda := getDB()[key]
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(agenda)
}

func delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])
	slice := getDB()
	copy(slice[key:], slice[key+1:])
	slice = slice[:len(slice)-1]
	updateData(slice)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

}
