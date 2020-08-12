package routes

import (
"api/controllers"
"net/http"
)

var userRoutes = []Route{
	Route{
		Uri: "/login",
		Method: http.MethodGet,
		Handler: controllers.LoginUser,
	},
}
