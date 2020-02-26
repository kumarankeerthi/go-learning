package router

import (
	"github.com/gorilla/mux"
	"github.com/kumarankeerthi/go-learning/whatsapp/controller"
)

func MessageRouter(r *mux.Router) *mux.Router {
	r.HandleFunc("/getMessage", controller.GetMessage).Methods("GET")
	r.HandleFunc("/sendMessage", controller.SendMessage).Methods("POST")
	return r
}
