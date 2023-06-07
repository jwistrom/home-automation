package main

import (
	"github.com/jwistrom/home-automation/tv-background/types"
	"log"
)

type Hardware struct {
	state types.SwitchState
	mode  types.TvBackgroundLightMode
	speed int
}

func NewHardware() *Hardware {
	return &Hardware{
		state: false,
		mode:  0,
		speed: 2,
	}
}

func (hardware *Hardware) GetStatus() types.Status {
	return types.Status{
		State: hardware.state,
		Mode:  hardware.mode,
		Speed: hardware.speed,
	}
}

func (hardware *Hardware) SwitchState() types.SwitchState {
	if hardware.state == types.On {
		hardware.state = types.Off
	} else {
		hardware.state = types.On
	}

	log.Printf("Switching state to %v\n", hardware.state)

	return hardware.state
}

func (hardware *Hardware) SetMode(mode types.TvBackgroundLightMode) {
	log.Printf("Setting mode to %v\n", mode)
	hardware.mode = mode
}

func (hardware *Hardware) SetSpeed(speed int) {
	log.Printf("Setting speed to %v\n", speed)
	hardware.speed = speed
}
