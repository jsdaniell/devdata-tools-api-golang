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
		Uri: "/suites",
		Method: http.MethodPost,
		Handler: controllers.CreateNewSuite,
	},
	Route{
		Uri: "/suites/{type}/{id}",
		Method: http.MethodDelete,
		Handler: controllers.DeleteSuite,
	},
	Route{
		Uri: "/suites/{type}/{id}/add",
		Method: http.MethodPost,
		Handler: controllers.AddNewItemOnSuite,
	},
	Route{
		Uri: "/suites/{type}/{id}",
		Method: http.MethodGet,
		Handler: controllers.GetAllItemsFromSuite,
	},
	Route{
		Uri: "/suites/{type}/{id}/{itemId}",
		Method: http.MethodPut,
		Handler: controllers.EditItemFromSuite,
	},
	Route{
		Uri: "/suites/{type}/{id}/{itemId}",
		Method: http.MethodDelete,
		Handler: controllers.DeleteItemFromSuite,
	},
}