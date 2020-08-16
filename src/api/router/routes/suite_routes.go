package routes

import (
	"github.com/jsdaniell/devdata-tools-api-golang/api/controllers"
	"net/http"
)

var suiteRoutes = []Route{
	Route{
		Uri: "/suites/{type}",
		Method: http.MethodGet,
		Handler: controllers.GetAllSuitesOfAType,

	},
	Route{
		Uri: "/suites/{type}/{name}",
		Method: http.MethodPost,
		Handler: controllers.CreateNewSuite,
	},
}