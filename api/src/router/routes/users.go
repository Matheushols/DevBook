package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Routes{
	{
		URI:                    "/users",
		Metode:                 http.MethodPost,
		Function:               controllers.CreateUser,
		RequiredAuthentication: false,
	},
	{
		URI:                    "/users",
		Metode:                 http.MethodGet,
		Function:               controllers.ListUsers,
		RequiredAuthentication: false,
	},
	{
		URI:                    "/users/{userId}",
		Metode:                 http.MethodGet,
		Function:               controllers.SearchUser,
		RequiredAuthentication: false,
	},
	{
		URI:                    "/users/{userId}",
		Metode:                 http.MethodPut,
		Function:               controllers.EditUser,
		RequiredAuthentication: false,
	},
	{
		URI:                    "/users/{userId}",
		Metode:                 http.MethodDelete,
		Function:               controllers.DeleteUser,
		RequiredAuthentication: false,
	},
}
