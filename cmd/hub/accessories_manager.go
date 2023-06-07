package main

import tvBackground "github.com/jwistrom/home-automation/tv-background/client"
import tvBackgroundTypes "github.com/jwistrom/home-automation/tv-background/types"

type AccessoriesStatus map[string]interface{}

type AccessoriesManager struct {
	tvBackgroundLightClient *tvBackground.TvBackgroundLightClient
}

func NewAccessoriesManager() *AccessoriesManager {
	light := tvBackground.NewTvBackgroundLightClient("localhost", 8081)
	light.SwitchState()
	light.SetMode(tvBackgroundTypes.ChaseTail)
	return &AccessoriesManager{light}
}

func (am *AccessoriesManager) getStatusOfAllAccessories() AccessoriesStatus {
	tvBackgroundStatus, _ := am.tvBackgroundLightClient.GetStatus()
	statuses := map[string]interface{}{
		"tv_background_light": tvBackgroundStatus,
	}

	return statuses
}
