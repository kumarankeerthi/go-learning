package router

import (
	"github.com/gorilla/mux"
	"github.com/kumarankeerthi/go-learning/rest-api-mongo/employee/controller"
)

func EmployeeRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/listEmployees", controller.GetAllEmployees).Methods("GET")
	router.HandleFunc("/employee/{id}", controller.GetEmployee).Methods("GET")
	return router
}
