package routes

import (
	"api/controllers"
	"net/http"
)

var testCasesRoutes = []Route{
	Route{
		Uri: "/tests-suites",
		Method: http.MethodGet,
		Handler: controllers.GetAllTestsSuites,
	},
	Route{
		Uri: "/tests-suites",
		Method: http.MethodDelete,
		Handler: controllers.DeleteTestSuite,
	},
	Route{
		Uri: "/tests-suites",
		Method: http.MethodPost,
		Handler: controllers.CreateTestSuite,
	},
	Route{
		Uri: "/tests-suites",
		Method: http.MethodPut,
		Handler: controllers.UpdateTestSuite,
	},
	Route{
		Uri: "/tests",
		Method: http.MethodGet,
		Handler: controllers.GetTestsFromATestSuite,
	},
	Route{
		Uri: "/tests/{id}",
		Method: http.MethodDelete,
		Handler: controllers.DeleteOneTestFromTestSuite,
	},
	Route{
		Uri: "/tests/{id}",
		Method: http.MethodPut,
		Handler: controllers.UpdateOneTestFromTestSuite,
	},
	Route{
		Uri: "/tests",
		Method: http.MethodPost,
		Handler: controllers.CreateOneTestOnTestSuite,
	},
}