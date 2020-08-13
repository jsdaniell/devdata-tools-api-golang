package routes

import (
	"github.com/jsdaniell/devdata-tools-api-golang/api/controllers"
	"net/http"
)

var userRoutes = []Route{
	Route{
		Uri:     "/login",
		Method:  http.MethodPost,
		Handler: controllers.LoginUser,
	},
}
