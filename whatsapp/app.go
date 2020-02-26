package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/kumarankeerthi/go-learning/whatsapp/data"
	"github.com/kumarankeerthi/go-learning/whatsapp/router"
)

type App struct {
	Router *mux.Router
	DB     *mongo.Database
}

func (app *App) InitializeApp() {

	ctx := context.Background()

	app.DB = data.InitDB(ctx)
	app.Router = router.InitRouter()

}

func (app *App) Run(add string) {
	srv := &http.Server{
		Handler:      app.Router,
		Addr:         add,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
