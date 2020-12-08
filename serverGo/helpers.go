package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func saveFile(r *http.Request, fileName string) (bool, error, string) {
	log.Println("Start Processing file uploading. Received File name:" + fileName)

	f1, handler, err := r.FormFile(fileName)
	if err != nil {
		log.Println("Error1")
		return false, err, ""
	}
	defer f1.Close()

	path := "./output/" + handler.Filename
	f2, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println("Error2")
		return false, err, ""
	}
	defer f2.Close()

	// Copy file to own path
	_, err = io.Copy(f2, f1)
	if err != nil {
		log.Println("Error3")
		return false, err, ""
	}

	return true, nil, handler.Filename
}

func randomParamGen() string {
	gen := rand.New(rand.NewSource(time.Now().UnixNano()))
	return "random=" + strconv.Itoa(gen.Intn(3000000))
}
