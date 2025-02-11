package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// The Routes represent all of the routes in the API
type Routes struct {
	URI                    string
	Metode                 string
	Function               func(http.ResponseWriter, *http.Request)
	RequiredAuthentication bool
}

func Configuration(r *mux.Router) *mux.Router {
	routes := usersRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Metode)
	}

	return r
}
