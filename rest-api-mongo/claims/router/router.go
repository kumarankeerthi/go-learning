package router

import (
	"github.com/gorilla/mux"
)

func InitializeRouters() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	router = ClaimsRouter(router)
	return router
}
