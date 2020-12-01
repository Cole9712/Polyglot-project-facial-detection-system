package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/julienschmidt/httprouter"
)

type Response struct {
	Success bool    `json:"success"`
	Error   *string `json:"error"`
	Url     string  `json:"url"`
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

		// var returnURL string = "http://localhost:8081/"
		fs := http.FileServer(http.Dir("/output/tempPic.jpg"))
		log.Fatal(http.ListenAndServe(":8081", fs))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// err = json.NewEncoder(w).Encode(Response{returnURL})
		// if err != nil {
		// 	panic(err)
		// }

	}
}

func fileUploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.ParseMultipartForm(32 << 20)
	log.Println("received request for uploading image")

	flag, err, filePath := saveFile(r, "uploadFileObj")
	if !flag {
		panic(err)
	}

	// invoke python program via Command line execution
	cmd := exec.Command("python3", "main.py", filePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatal("cmd Run failed")
		panic(err)
	}

	// prevernt CORS issue
	w.Header().Set("Access-Control-Allow-Methods", " GET, POST, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token")
	response := Response{true, nil, "http://0.0.0.0:8082/" + filePath}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)

	// w.WriteHeader(http.StatusCreated)

}
