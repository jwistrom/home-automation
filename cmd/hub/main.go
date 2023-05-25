package main

import (
	"html/template"
	"log"
	"net/http"
)

var htmlTemplate *template.Template

func main() {
	loadTemplate()
	http.HandleFunc("/", handleRoot)

	log.Println("Serving on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func loadTemplate() {
	var err error
	htmlTemplate, err = template.ParseFiles("index.html")
	if err != nil {
		log.Fatalf("Load template error %s", err)
	}
}

func handleRoot(w http.ResponseWriter, req *http.Request) {
	err := htmlTemplate.Execute(w, nil)
	if err != nil {
		log.Printf("Template render error: %s", err)
		http.Error(w, err.Error(), 500)
	}
}
