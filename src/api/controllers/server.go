package controllers

import "net/http"

func ServerRunning(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server Running..."))
}