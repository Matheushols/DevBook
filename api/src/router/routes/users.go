package routes

import "net/http"

var usersRoutes = []Routes{
	{
		URI:    "/users",
		Metode: http.MethodPost,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},
		RequiredAuthentication: false,
	},
	{
		URI:    "/users",
		Metode: http.MethodGet,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},
		RequiredAuthentication: false,
	},
	{
		URI:    "/users/{userId}",
		Metode: http.MethodGet,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},
		RequiredAuthentication: false,
	},
	{
		URI:    "/users/{userId}",
		Metode: http.MethodPut,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},
		RequiredAuthentication: false,
	},
	{
		URI:    "/users/{userId}",
		Metode: http.MethodDelete,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},
		RequiredAuthentication: false,
	},
}
