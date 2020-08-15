package middlewares

import (
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

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}