package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type PageVariables struct {
	AsciiArt string
}

func main() {
	http.HandleFunc("/", DisplayPage)
	http.HandleFunc("/ascii-art", GenerateAsciiArtHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // ローカル用デフォルト
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func DisplayPage(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Print("template parsing error: ", err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found (404)"))
		return
	}

	var pageVariables PageVariables
	err = t.Execute(w, pageVariables)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}

func GenerateAsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	text := r.Form.Get("text")
	renderingType := r.Form.Get("renderingType")

	if text == "" || renderingType == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request: Invalid parameters (400)"))
		return
	}

	asciiArt := GenerateAsciiArt(text, renderingType)

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found (404)"))
		return
	}

	var pageVariables PageVariables
	pageVariables.AsciiArt = asciiArt

	err = t.Execute(w, pageVariables)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error (500)"))
		return
	}

	w.WriteHeader(http.StatusOK)
}
