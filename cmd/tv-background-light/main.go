package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var hardware *Hardware

func main() {

	hardware = NewHardware()

	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/state", handleState)
	http.HandleFunc("/mode", handleMode)
	http.HandleFunc("/speed", handleSpeed)
	http.HandleFunc("/color", handleColor)

	log.Println("Serving on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func handleRoot(w http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		jsonResponse, err := json.Marshal(hardware.GetStatus())
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		log.Println("Sending hardware status")
		_, err = w.Write(jsonResponse)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
	}
}

func handleState(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPut {
		newState := hardware.SwitchState()
		_, err := w.Write([]byte(strconv.FormatBool(bool(newState))))
		if err != nil {
			log.Println("Failed to write to response")
		}
	}
}

func handleMode(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPut {
		params := req.URL.Query()
		if params.Has("mode") {
			inputMode := params.Get("mode")
			intMode, err := strconv.Atoi(inputMode)
			if err != nil {
				log.Printf("Failed to parse mode %s\n", inputMode)
				return
			}

			hardware.SetMode(intMode)
			_, _ = w.Write([]byte(fmt.Sprintf("%d", intMode)))
		}
	}
}

func handleSpeed(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPut {
		params := req.URL.Query()
		if params.Has("speed") {
			inputSpeed := params.Get("speed")
			intSpeed, err := strconv.Atoi(inputSpeed)
			if err != nil {
				log.Printf("Failed to parse speed %s\n", inputSpeed)
				return
			}

			hardware.SetSpeed(intSpeed)
			_, _ = w.Write([]byte(fmt.Sprintf("%d", intSpeed)))
		}
	}
}

func handleColor(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPut {
		params := req.URL.Query()
		if params.Has("color") {
			inputColor := params.Get("color")

			hardware.SetColor(inputColor)
			_, _ = w.Write([]byte(fmt.Sprintf("%s", inputColor)))
		}
	}
}
