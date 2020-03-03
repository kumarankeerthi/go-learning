package api

import (
	"fmt"
	"net/http"

	"github.com/kumarankeerthi/go-learning/graphql/core"
)

type RequestHanlder interface {
	SavePerson(w http.ResponseWriter, r *http.Request)
	GetPersonById(w http.ResponseWriter, r *http.Request)
}

type handle struct {
	personServ core.PersonService
}

func CreateHandler(ps core.PersonService) RequestHanlder {
	return &handle{personServ: ps}
}

func (h *handle) SavePerson(w http.ResponseWriter, r *http.Request) {

}

func (h *handle) GetPersonById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in api layer")
	h.personServ.GetPersonById("1234")
}
