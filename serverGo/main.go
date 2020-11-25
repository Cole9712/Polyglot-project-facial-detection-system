package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	// fileRouter := httprouter.New()
	// router.GET("/", indexHandler(router))
	router.POST("/uploadMine", fileUploadHandler)
	router.ServeFiles("/*filepath", http.Dir("./output"))

	log.Fatal(http.ListenAndServe(":8082", router))
	// log.Fatal(http.ListenAndServe(":8081", fileRouter))
}
