package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/gorilla/mux"
	"github.com/kumarankeerthi/go-learning/graphql/api"
	"github.com/kumarankeerthi/go-learning/graphql/core"
	"github.com/kumarankeerthi/go-learning/graphql/data"
)

func main() {

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "person",
			"time", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller)

	}

	level.Info(logger).Log("msg", "GraphQl Person service started")
	defer level.Info(logger).Log("msg", "GraphQl Person service stopped")

	personRepo := data.CreateRepository("localhost:27017", "graphql")

	personService := core.CreatePersonService(personRepo, logger)

	router := mux.NewRouter()

	//**** below is for Endpoint IMplementation using Go-Kit
	addPersonEndpointHanlder := httptransport.NewServer(
		api.AddPersonEndpoint(personService),
		api.DecodeAddPersonRequest,
		api.EncodeAddPersonResponse,
	)

	getPersonEndpointHandler := httptransport.NewServer(
		api.GetPersonEndpoint(personService),
		api.DecodeGetPersonRequest,
		api.EncodeGetPersonResponse,
	)
	router.Methods("POST").Handler(addPersonEndpointHanlder).Path("/addPerson")
	router.Methods("GET").Handler(getPersonEndpointHandler).Path("/getPerson/{PersonId}")
	//**** Below is for straight fwd implementation of router

	// h := api.CreateHandler(personService)
	// router.HandleFunc("/getPerson/{PersonId}", h.FetchPersonById).Methods("GET")
	// router.HandleFunc("/addPerson", h.AddPerson).Methods("POST")

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		fmt.Println("Listening on port 8000")
		errs <- srv.ListenAndServe()
	}()
	level.Error(logger).Log("exit", <-errs)
}
