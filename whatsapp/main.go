package main

import (
	"log"
	"net/http"
	"time"

	"github.com/kumarankeerthi/go-learning/webserver/router"
)

func main() {
	r := router.InitRouter()
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
