package main

import "github.com/jwistrom/home-automation/cmd/hub/accessories"

type AccessoriesStatus map[string]interface{}

type AccessoriesManager struct {
	tvBackgroundLight *accessories.TvBackgroundLight
}

func NewAccessoriesManager() *AccessoriesManager {
	light := &accessories.TvBackgroundLight{}
	light.SwitchState()
	light.SetMode(accessories.ChaseTail)
	return &AccessoriesManager{light}
}

func (am *AccessoriesManager) getStatusOfAllAccessories() AccessoriesStatus {
	statuses := map[string]interface{}{
		am.tvBackgroundLight.GetId(): am.tvBackgroundLight.GetStatus(),
	}

	return statuses
}
