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
	router.POST("/uploadMinePost", fileUploadHandler)
	router.POST("/uploadMinePost/swap", swapRequestHandler)
	router.POST("/home", homeFacialHandler)
	router.ServeFiles("/*filepath", http.Dir("./facialOutput"))
	// router.ServeFiles("/homeOutput/*filepath", http.Dir("./homeOutput"))
	corsHandler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe("0.0.0.0:8082", corsHandler))
}
