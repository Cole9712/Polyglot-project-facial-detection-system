package main

import (
	"log"
	"net/http"
	"encoding/json"

	"github.com/julienschmidt/httprouter"
)

type Response struct {
	ImageURL string `json:"ImageURL"`
}

func indexHandler(router *httprouter.Router) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		img, err := parseImageJSON(r, w)
		if err != nil {
			panic("Error parseImageJSON")
		}
		message := saveBase64(img.ImageBase64, img.ImageType)
		if message != "success!" {
			panic("save Base64 to file unsuccessfully")
		}
		
		var returnURL string = "http://localhost:8081/"
		fs := http.FileServer(http.Dir("/output/tempPic.jpg"))
		log.Fatal(http.ListenAndServe(":8081", fs))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(Response{returnURL})
		if err != nil {
			panic(err)
		}

	}

}