package routes

import "net/http"

// The Routes represent all of the routes in the API
type Routes struct {
	URI                    string
	Metode                 string
	Function               func(http.ResponseWriter, *http.Request)
	RequiredAuthentication bool
}
