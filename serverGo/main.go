package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func main() {
	
	router := httprouter.New()
	log.Print("REST Server Started")
	// fileRouter := httprouter.New()
	// router.GET("/", indexHandler(router))
	router.POST("/uploadMine", fileUploadHandler)
	router.ServeFiles("/*filepath", http.Dir("./facialOutput"))
	corsHandler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe("0.0.0.0:8082", corsHandler))
}
