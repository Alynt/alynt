package app

import (
	"log"
	"net/http"

	"github.com/Alynt/alynt/config"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	Server *http.Server
}

func (a *App) Initialize(c *config.Config) {
	a.Router = mux.NewRouter()
	a.Server = &http.Server{
		Addr:    c.ServerPort,
		Handler: a.Router,
	}

	a.setRouting()
}

func (a *App) Run() {
	log.Println("Starting in ->", a.Server.Addr)

	if err := a.Server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
