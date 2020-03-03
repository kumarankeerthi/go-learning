package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
	"github.com/kumarankeerthi/go-learning/graphql/core"
)

type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)

// type PersonRequest struct {
// 	FirstName string `json:"firstName"`
// 	LastName  string `json:"lastName"`
// 	Age       int    `json:"age"`
// }

type PersonResponse struct {
	PID string `json:PersonId"`
}

func getPersonEndpoint(svc core.PersonService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(PersonResponse)
		person := svc.FetchPersonById(req.PID)
		return person, nil
	}
}

func GetPersonEndpoint(svc core.PersonService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		res := request.(PersonResponse)
		p := svc.FetchPersonById(res.PID)
		return p, nil
	}
}

func DecodeGetPersonRequest(_ context.Context, r *http.Request) (interface{}, error) {
	pid := mux.Vars(r)["PersonId"]
	return PersonResponse{PID: pid}, nil
}

func EncodeGetPersonResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func AddPersonEndpoint(svc core.PersonService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		res := request.(core.Person)
		result, err := svc.AddPerson(res)
		if err != nil {
			return PersonResponse{PID: ""}, err
		}
		return PersonResponse{PID: result}, nil

	}
}

func DecodeAddPersonRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var p core.Person
	err := json.NewDecoder(r.Body).Decode(&p)
	return p, err

}

func EncodeAddPersonResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
