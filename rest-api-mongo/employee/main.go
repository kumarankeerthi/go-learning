package main

import (
	"log"
	"net/http"

	"github.com/kumarankeerthi/go-learning/rest-api-mongo/employee/router"
)

func main() {
	log.Printf("Starting Employee microservice!!")
	router := router.InitRouters()

	log.Fatal(http.ListenAndServe(":8000", router))
}
