package routes

import (
	"api/src/controllers"
	"net/http"
)

var loginRoute = Routes{
	URI:                    "/login",
	Metode:                 http.MethodPost,
	Function:               controllers.Login,
	RequiredAuthentication: false,
}
