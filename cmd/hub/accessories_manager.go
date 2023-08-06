package main

import (
	tvBackground "github.com/jwistrom/home-automation/tv-background/client"
	"log"
)

type AccessoriesStatus map[string]interface{}

type AccessoriesManager struct {
	tvBackgroundLightClient *tvBackground.TvBackgroundLightClient
}

func NewAccessoriesManager(conf accessoriesManagerConf) *AccessoriesManager {
	lightClient := tvBackground.NewTvBackgroundLightClient(conf.tvBackgroundIp, conf.tvBackgroundPort)
	return &AccessoriesManager{lightClient}
}

func (am *AccessoriesManager) getStatusOfAllAccessories() AccessoriesStatus {
	tvBackgroundStatus, err := am.tvBackgroundLightClient.GetStatus()
	if err != nil {
		log.Println("Failed to get status of tv background light")
	}
	statuses := map[string]interface{}{
		"tv_background_light": tvBackgroundStatus,
	}

	return statuses
}

type accessoriesManagerConf struct {
	tvBackgroundIp   string
	tvBackgroundPort int
}
