package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/kumarankeerthi/go-learning/hexa-design-webapp/api"
	"github.com/kumarankeerthi/go-learning/hexa-design-webapp/chat"
	"github.com/kumarankeerthi/go-learning/hexa-design-webapp/repository"
)

func main() {
	fmt.Println("Hexagonal Design Pattern implementation")

	chatRepo := repository.NewConnection("localhost:27017", "local")
	chatServ := chat.NewChatService(chatRepo)
	h := api.CreateHandler(chatServ)

	router := mux.NewRouter()
	router.HandleFunc("/viewChat", h.ViewChat).Methods("GET")
	router.HandleFunc("/sendMessage", h.SendMessage).Methods("POST")

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
