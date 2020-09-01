package api

import (
	"fmt"
	"github.com/jsdaniell/devdata-tools-api-golang/api/router"
	"github.com/jsdaniell/devdata-tools-api-golang/config"
	"log"
	"net/http"
)

func Run() {
	config.Load()
	fmt.Printf("\nListening... localhost:%d\n", config.PORT)
	listen(config.PORT)
}

func listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
