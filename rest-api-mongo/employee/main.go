package main

import (
	"log"
	"net/http"

	"github.com/kumarankeerthi/go-learning/rest-api-mongo/employee/data"
	"github.com/kumarankeerthi/go-learning/rest-api-mongo/employee/model"
	"github.com/kumarankeerthi/go-learning/rest-api-mongo/employee/router"
	"github.com/mongodb/mongo-go-driver/bson"
)

func main() {
	log.Printf("Starting Employee microservice!!")
	router := router.InitRouters()
	collection := data.CollectionHandle("local", "Employee")
	result := data.FindEmployeeByID("asd", collection)
	kk := bson.E(result, model.Employee)
	log.Print(kk)
	log.Fatal(http.ListenAndServe(":8000", router))
}
