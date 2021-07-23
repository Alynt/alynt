package app

import (
	"github.com/Alynt/alynt/app/routes"
)

func (a *App) setRouting() {
	api := a.Router.PathPrefix("/api").Subrouter()

	// Routes
	api.HandleFunc("/test", routes.Test)
}
