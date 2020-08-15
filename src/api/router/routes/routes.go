package routes

import (
	"github.com/gorilla/mux"
	"github.com/jsdaniell/devdata-tools-api-golang/api/middlewares"
	"net/http"
)

type Route struct {
	Uri           string
	Method        string
	Handler       func(http.ResponseWriter, *http.Request)
	Open bool
}

func Load() []Route {

	routes := [][]Route{
		serverRoutes,
		userRoutes,
		suiteRoutes,
	}

	var joinedRoutes []Route

	for _, r := range routes {
		joinedRoutes = append(joinedRoutes, r...)
	}

	return joinedRoutes
}

func SetupRoutes(r *mux.Router) *mux.Router {

	for _, route := range Load() {
		r.HandleFunc(route.Uri, route.Handler).Methods(route.Method)
	}

	return r
}

func SetupRoutesWithMiddlewares(r *mux.Router) *mux.Router {

	for _, route := range Load() {
		r.HandleFunc(route.Uri, middlewares.SetMiddlewareLogger(
			middlewares.SetMiddlewareJSON(route.Handler, route.Open))).Methods(route.Method, "OPTIONS")
	}

	return r
}
