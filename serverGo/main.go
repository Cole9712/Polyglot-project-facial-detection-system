package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	log.Print("REST Server Started")
	// fileRouter := httprouter.New()
	// router.GET("/", indexHandler(router))
	router.POST("/uploadMine", fileUploadHandler)
	router.ServeFiles("/*filepath", http.Dir("./facialOutput"))
	
	log.Fatal(http.ListenAndServe(":8082", router))
	// log.Fatal(http.ListenAndServe(":8081", fileRouter))
}
