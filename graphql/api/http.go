package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kumarankeerthi/go-learning/graphql/core"
)

type RequestHanlder interface {
	AddPerson(w http.ResponseWriter, r *http.Request)
	FetchPersonById(w http.ResponseWriter, r *http.Request)
}

type handle struct {
	personServ core.PersonService
}

func CreateHandler(ps core.PersonService) RequestHanlder {
	return &handle{personServ: ps}
}

func (h *handle) AddPerson(w http.ResponseWriter, r *http.Request) {
	var p core.Person
	json.NewDecoder(r.Body).Decode(&p)
	result, err := h.personServ.AddPerson(p)
	if err != nil {
		fmt.Println("error adding person", p)
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintln("Sucessfully added person :", result)))
}

func (h *handle) FetchPersonById(w http.ResponseWriter, r *http.Request) {
	pid := mux.Vars(r)["PersonId"]
	p := h.personServ.FetchPersonById(pid)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&p)

}
