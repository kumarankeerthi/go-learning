package main

import (
	"log"
	"net/http"

	"github.com/kumarankeerthi/go-learning/rest-api-mongo/claims/router"
)

func main() {
	router := router.InitializeRouters()

	log.Fatal(http.ListenAndServe(":8080", router))
}
