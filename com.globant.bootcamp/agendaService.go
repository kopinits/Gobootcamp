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
	w.Header().Set("Content-Type", "application/json")
	for _, agenda := range ListAll() {
		json.NewEncoder(w).Encode(agenda)
	}
	w.WriteHeader(http.StatusAccepted)
}

func saveAgenda(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var agenda Agenda
	json.Unmarshal(requestBody, &agenda)
	Save(agenda)
	doReturn(w)
}

func getById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])
	if len(getDatabase()) >= key {
		agenda := GetById(key)
		doReturn(w)
		json.NewEncoder(w).Encode(agenda)
	} else {
		json.NewEncoder(w).Encode("Agenda id " + vars["id"] + " not found")
	}
}

func delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])
	if Delete(key) {
		doReturn(w)
	} else {
		json.NewEncoder(w).Encode("Agenda id " + vars["id"] + " not found")
	}
}

func doReturn(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
}
