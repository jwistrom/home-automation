package main

import (
	"github.com/SimonWaldherr/ws2812/pixarray"
	"github.com/jwistrom/home-automation/tv-background/types"
	"log"
)

const (
	ledCount   = 100
	colorCount = 3
	ledPin     = 18
	defaultDma = 10
	frequency  = 800000
)

type Hardware struct {
	state types.SwitchState
	mode  int
	speed int
	color string
	pa    *pixarray.PixArray
}

func NewHardware() *Hardware {
	order := pixarray.StringOrders["GRB"]
	leds, err := pixarray.NewWS281x(ledCount, colorCount, order, frequency, defaultDma, []int{ledPin})
	if err != nil {
		log.Fatalf("Failed creating WS281x: %v", err)
	}

	pa := pixarray.NewPixArray(ledCount, colorCount, leds)

	var p pixarray.Pixel

	p.R = 0
	p.G = 0
	p.B = 0
	pa.SetAll(p)
	pa.Write()

	return &Hardware{
		state: false,
		mode:  0,
		speed: 2,
		color: "000000FF",
		pa:    pa,
	}
}

func (hardware *Hardware) GetStatus() types.Status {
	return types.Status{
		State: hardware.state,
		Mode:  hardware.mode,
		Speed: hardware.speed,
		Color: hardware.color,
	}
}

func (hardware *Hardware) SwitchState() types.SwitchState {
	if hardware.state == types.On {
		hardware.switchOff()
	} else {
		hardware.switchOn()
	}

	log.Printf("Switching state to %v\n", hardware.state)

	return hardware.state
}

func (hardware *Hardware) SetMode(mode int) {
	log.Printf("Setting mode to %v\n", mode)
	hardware.mode = mode
}

func (hardware *Hardware) SetSpeed(speed int) {
	log.Printf("Setting speed to %v\n", speed)
	hardware.speed = speed
}

func (hardware *Hardware) SetColor(color string) {
	log.Printf("Setting color to %s", color)
	hardware.color = color
}

func (hardware *Hardware) switchOn() {
	mode := TvBackgroundLightModes[hardware.mode]
	mode.switchOn(hardware)
	hardware.state = types.On
}

func (hardware *Hardware) switchOff() {
	mode := TvBackgroundLightModes[hardware.mode]
	mode.switchOff(hardware)
	hardware.state = types.Off
}
