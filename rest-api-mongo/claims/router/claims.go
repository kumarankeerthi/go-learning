package router

import (
	"github.com/gorilla/mux"
	"github.com/kumarankeerthi/go-learning/rest-api-mongo/claims/controller"
)

func ClaimsRouter(router *mux.Router) *mux.Router {
	router.HandleFunc("/claims", controller.GetClaims).Methods("GET")
	router.HandleFunc("/claims/{id}", controller.GetClaimsOfEmployee).Methods("GET")
	return router
}
