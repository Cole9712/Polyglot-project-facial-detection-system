package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type imageGot struct {
	ImageType   string `json:"ImageType"`
	ImageBase64 string `json:"ImageBase64"`
}

func saveFile(r *http.Request, fileName string) (bool, error, string) {
	log.Println("Start Processing file uploading. Received File name:" + fileName)
	
	f1, handler, err := r.FormFile(fileName)
	if err != nil {
		log.Println("Error1")
		return false, err, ""
	}
	defer f1.Close()

	path := "./output/"+handler.Filename
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

func parseImageJSON(r *http.Request, w http.ResponseWriter) (imageGot, error) {
	// declare a received image struct type
	var receivedImage imageGot

	err := json.NewDecoder(r.Body).Decode(&receivedImage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

	}

	return receivedImage, err
}

// Save a JPG to ./output/
func saveJPG(input image.Image) string {
	fileName := "/output/tempPic.jpg"
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
		return "error: saveJPG openFile"
	}

	err = jpeg.Encode(f, input, &jpeg.Options{Quality: 100})
	if err != nil {
		log.Fatal(err)
		return "error: saveJPG jpeg.Encode"
	}

	fmt.Println("JPG File created!")
	return "success!"
}

// adapted from https://github.com/rravishankar/golangtraining/tree/master/test/jpegencode
// convert base64 string to image file (JPG)
func saveBase64(input string, format string) string {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(input))
	i, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
		return "error: saveBase64"
	}

	message := saveJPG(i)

	return message
}
