package main

import (
	"log"
	"net/http"

	"github.com/kumarankeerthi/go-learning/rest-api-mongo/employee/data"
	"github.com/kumarankeerthi/go-learning/rest-api-mongo/employee/router"
)

func main() {
	log.Printf("Starting Employee microservice!!")
	router := router.InitRouters()
	collection := data.CollectionHandle("local", "Employee")
	result := data.FindEmployeeByID("asd", collection)
	log.Print(result)
	log.Fatal(http.ListenAndServe(":8000", router))
}
