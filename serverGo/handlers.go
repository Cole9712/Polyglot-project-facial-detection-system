package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type Response struct {
	Success bool    `json:"success"`
	Error   *string `json:"error"`
	Url     string  `json:"url"`
}

type SwapRequest struct {
	File1 string `json:"file1"`
	File2 string `json:"file2"`
}

type HomeFacialRequest struct {
	ImgUrl string `json:"imgUrl"`
	ID     int    `json:"id"`
	Ext    string `json:"ext"`
}

func homeFacialHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Print("Received Facial Detection Request for Home Page.")
	// save file from request
	var hfr HomeFacialRequest
	err := json.NewDecoder(r.Body).Decode(&hfr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response, e := http.Get(hfr.ImgUrl)
	if e != nil {
		log.Fatal(e)
	}
	defer response.Body.Close()
	path := "./homeOutput/" + strconv.Itoa(hfr.ID) + "." + hfr.Ext
	f1, _ := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	defer f1.Close()

	_, err = io.Copy(f1, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Save File Successfully!")
	cmd := exec.Command("python3", "main.py", "detectAPI", strconv.Itoa(hfr.ID)+"."+hfr.Ext)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatal("cmd Run failed")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	returnResponse := Response{true, nil, "http://127.0.0.1:8082/" + strconv.Itoa(hfr.ID) + "." + hfr.Ext + "?" + randomParamGen()}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(returnResponse)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	log.Println("Response sent with StatusOK")

}

func swapRequestHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var sr SwapRequest
	err := json.NewDecoder(r.Body).Decode(&sr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	cmd := exec.Command("python3", "main.py", "swap", sr.File1, sr.File2)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatal("cmd Run failed")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := Response{true, nil, "http://127.0.0.1:8082/" + sr.File1 + "?" + randomParamGen()}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	log.Println("Response sent with StatusOK")

}

func fileUploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.ParseMultipartForm(32 << 20)
	log.Println("received request for uploading image")

	flag, err, filePath := saveFile(r, "uploadFileObj")
	if !flag {
		panic(err)
	}

	// invoke python program via Command line execution
	cmd := exec.Command("python3", "main.py", "detect", filePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatal("cmd Run failed")
		panic(err)
	}

	// prevernt CORS issue
	// w.Header().Set("Access-Control-Allow-Methods", " GET, POST, PATCH, PUT, DELETE, OPTIONS")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token")
	response := Response{true, nil, "http://127.0.0.1:8082/" + filePath + "?" + randomParamGen()}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)

	// w.WriteHeader(http.StatusCreated)

}
