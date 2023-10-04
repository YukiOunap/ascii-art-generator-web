package main

import (
	"html/template"
	"log"
	"net/http"
)

type PageVariables struct {
	AsciiArt string
}

func main() {
	http.HandleFunc("/", DisplayPage)
	http.HandleFunc("/ascii-art", GenerateAsciiArtHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
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

	var pageVariables PageVariables
	pageVariables.AsciiArt = asciiArt

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Print("template parsing error: ", err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found (404)"))
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error (500)"))
		return
	}

	err = t.Execute(w, pageVariables)
	if err != nil {
		log.Print("template executing error: ", err)
	}

	w.WriteHeader(http.StatusOK)
}
