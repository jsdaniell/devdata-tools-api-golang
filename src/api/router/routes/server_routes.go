package routes

import (
	"github.com/jsdaniell/devdata-tools-api-golang/api/controllers"
	"net/http"
)

var serverRoutes = []Route{
	Route{
		Uri:     "/",
		Method:  http.MethodGet,
		Handler: controllers.ServerRunning,
		Open: true,
	},
}
