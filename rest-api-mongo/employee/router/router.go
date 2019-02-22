package router

import (
	"github.com/gorilla/mux"
)

func InitRouters() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	router = EmployeeRouters(router)
	return router
}
