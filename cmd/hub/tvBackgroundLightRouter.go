package main

import (
	"fmt"
	tvBackground "github.com/jwistrom/home-automation/tv-background/client"
	"log"
	"net/http"
	"strconv"
)

type TvBackgroundLightRouter struct {
	tvBackgroundLightClient *tvBackground.TvBackgroundLightClient
}

func (r *TvBackgroundLightRouter) HandleState(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPut {
		newState, _ := r.tvBackgroundLightClient.SwitchState()
		log.Printf("Switching Tv background light to %s\n", newState)
		_, _ = w.Write([]byte(fmt.Sprintf("%v", newState)))
	}
}

func (r *TvBackgroundLightRouter) HandleMode(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPut {
		params := req.URL.Query()
		if params.Has("mode") {
			inputMode := params.Get("mode")
			intMode, err := strconv.Atoi(inputMode)
			if err != nil {
				log.Printf("Failed to parse mode %s\n", inputMode)
				return
			}

			r.tvBackgroundLightClient.SetMode(intMode)
			log.Printf("Setting new Tv background light mode to %d\n", intMode)
			_, _ = w.Write([]byte(fmt.Sprintf("%v", intMode)))
		}
	}
}

func (r *TvBackgroundLightRouter) HandleSpeed(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPut {
		params := req.URL.Query()
		if params.Has("speed") {
			inputSpeed := params.Get("speed")
			intSpeed, err := strconv.Atoi(inputSpeed)
			if err != nil {
				log.Printf("Failed to parse speed %s\n", inputSpeed)
				return
			}

			r.tvBackgroundLightClient.SetSpeed(intSpeed)
			log.Printf("Setting new Tv background light speed to %d\n", intSpeed)
			_, _ = w.Write([]byte(fmt.Sprintf("%d", intSpeed)))
		}
	}
}

func (r *TvBackgroundLightRouter) HandleColor(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPut {
		params := req.URL.Query()
		if params.Has("color") {
			inputColor := params.Get("color")

			r.tvBackgroundLightClient.SetColor(inputColor)
			log.Printf("Setting new Tv background light speed to %s\n", inputColor)
			_, _ = w.Write([]byte(fmt.Sprintf("%s", inputColor)))
		}
	}
}
