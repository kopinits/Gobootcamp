package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
	data   []Agenda
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GoLand Bootcamp")
}

func (a *App) initializeRoutes() {
	a.Router = mux.NewRouter().StrictSlash(true)
	a.Router.HandleFunc("/", homePage).Methods("GET")
	a.Router.HandleFunc("/agendas", getAgendas).Methods("GET")
	a.Router.HandleFunc("/agenda", saveAgenda).Methods("POST")
	a.Router.HandleFunc("/agenda/{id}", getById).Methods("GET")
	a.Router.HandleFunc("/agenda/{id}", delete).Methods("DELETE")
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

var A = App{}

func getDB() []Agenda {
	return A.data
}

func updateData(data []Agenda) {
	A.data = data
}

func main() {
	A.data = make([]Agenda, 0)
	A.initializeRoutes()
	A.Run(":8874")
}
