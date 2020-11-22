package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", indexHandler(router))

	log.Fatal(http.ListenAndServe(":8082", router))
}
