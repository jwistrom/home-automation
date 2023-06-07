package main

import (
	"html/template"
	"log"
	"net/http"
)

var htmlTemplate *template.Template
var am *AccessoriesManager

func main() {
	am = NewAccessoriesManager()
	loadTemplate()

	http.HandleFunc("/", handleRoot)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	registerTvBackgroundLightRouter("/tvbackgroundlight")

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
	err := htmlTemplate.Execute(w, map[string]interface{}(am.getStatusOfAllAccessories()))
	if err != nil {
		log.Printf("Template render error: %s", err)
		http.Error(w, err.Error(), 500)
	}
}

func registerTvBackgroundLightRouter(prefix string) {
	router := &TvBackgroundLightRouter{am.tvBackgroundLightClient}
	http.HandleFunc(prefix+"/state", router.HandleState)
	http.HandleFunc(prefix+"/mode", router.HandleMode)
	http.HandleFunc(prefix+"/speed", router.HandleSpeed)
}
