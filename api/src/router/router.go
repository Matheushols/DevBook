package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// The Generate function will return a router with the configured routes.
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Configuration(r)
}
