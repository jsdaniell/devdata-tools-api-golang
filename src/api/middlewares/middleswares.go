package middlewares

import (
	"fmt"
	"github.com/jsdaniell/devdata-tools-api-golang/api/responses"
	"log"
	"net/http"
)

func SetMiddlewareLogger(next http.HandlerFunc) http.HandlerFunc{

	return func(w http.ResponseWriter, r *http.Request){

		// TODO: When in production setup to https://devdata.tools

		w.Header().Set("Access-Control-Allow-Origin", "*")

		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		log.Println("%s %s%s %s", r.Method, r.Host, r.RequestURI, r.Proto)

		if (*r).Method == "OPTIONS" {
			return
		}



		next(w, r)
	}
}

func SetMiddlewareJSON(next http.HandlerFunc, openRoute bool) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")

		if !openRoute {
			auth := r.Header.Get("Authorization")
			if auth == "" {
				responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("missing authorization token"))
				return
			}
		}

		next(w, r)
	}
}