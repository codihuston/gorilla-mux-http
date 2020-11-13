// app.go

package main

import (
	"github.com/codihuston/gorilla-mux-http/api/v1/controllers"
	"github.com/codihuston/gorilla-mux-http/db"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize(user, password, dbname string) {

	db.Initialize(user, password, dbname)

	a.Router = mux.NewRouter()

	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8010", a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	a.Router.HandleFunc("/product", controllers.CreateProduct).Methods("POST")
	a.Router.HandleFunc("/product/{id:[0-9]+}", controllers.GetProduct).Methods("GET")
	a.Router.HandleFunc("/product/{id:[0-9]+}", controllers.UpdateProduct).Methods("PUT")
	a.Router.HandleFunc("/product/{id:[0-9]+}", controllers.DeleteProduct).Methods("DELETE")
}
