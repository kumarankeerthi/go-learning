package router

import (
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(false)
	r = MessageRouter(r)
	return r
}
