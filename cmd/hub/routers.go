package main

import (
	"fmt"
	"github.com/jwistrom/home-automation/cmd/hub/accessories"
	"log"
	"net/http"
	"strconv"
)

type TvBackgroundLightRouter struct {
	tvBackgroundLight *accessories.TvBackgroundLight
}

func (r *TvBackgroundLightRouter) HandleState(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPut {
		newState := r.tvBackgroundLight.SwitchState()
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

			newMode := accessories.TvBackgroundLightMode(intMode)
			r.tvBackgroundLight.SetMode(newMode)
			log.Printf("Setting new Tv background light mode to %d\n", newMode)
			_, _ = w.Write([]byte(fmt.Sprintf("%v", newMode)))
		}
	}
}
