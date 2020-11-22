package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"os"
	"strings"
)

type imageGot struct {
	ImageType   string `json:"ImageType"`
	ImageBase64 string `json:"ImageBase64"`
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
