package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/kumarankeerthi/go-learning/graphql/api"
	"github.com/kumarankeerthi/go-learning/graphql/core"
	"github.com/kumarankeerthi/go-learning/graphql/data"
)

func main() {
	fmt.Println("GraphQL in Go!!")

	personRepo := data.CreateRepository("localhost:27017")

	personService := core.CreatePersonService(personRepo)

	h := api.CreateHandler(personService)

	router := mux.NewRouter()
	router.HandleFunc("/getPerson", h.GetPersonById).Methods("GET")

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
