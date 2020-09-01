package router

import (
	"github.com/gorilla/mux"
	"github.com/jsdaniell/devdata-tools-api-golang/api/router/routes"
)

func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	return routes.SetupRoutesWithMiddlewares(r)
}