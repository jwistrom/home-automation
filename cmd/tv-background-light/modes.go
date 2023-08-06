package main

import (
	colorconverter "github.com/SimonWaldherr/ColorConverterGo"
	"github.com/SimonWaldherr/ws2812/pixarray"
	"math"
	"strconv"
	"time"
)

const delayMs = 100

var TvBackgroundLightModes map[int]TvBackgroundLightMode = map[int]TvBackgroundLightMode{
	0: standard{},
	1: circle{},
	2: blink{},
	3: opposite{},
	4: chaseTail{},
	5: hug{},
	6: explosion{},
}

type TvBackgroundLightMode interface {
	switchOn(hardware *Hardware)
	switchOff(hardware *Hardware)
}

type standard struct{}
type circle struct{}
type blink struct{}
type opposite struct{}
type chaseTail struct{}
type hug struct{}
type explosion struct{}

func (s standard) switchOn(hardware *Hardware) {
	var p pixarray.Pixel
	p.R, p.G, p.B = toRgb(hardware.color)
	hardware.pa.SetAll(p)
	hardware.pa.Write()
}

func (s standard) switchOff(hardware *Hardware) {
	var p pixarray.Pixel
	p.R, p.G, p.B = 0, 0, 0
	hardware.pa.SetAll(p)
	hardware.pa.Write()
}

func (c circle) switchOn(hardware *Hardware) {
	var p pixarray.Pixel
	p.R, p.G, p.B = toRgb(hardware.color)
	for ledIdx := 0; ledIdx < hardware.pa.NumPixels(); ledIdx++ {
		hardware.pa.SetOne(ledIdx, p)
		hardware.pa.Write()
		time.Sleep(time.Millisecond * delayMs)
	}
}

func (c circle) switchOff(hardware *Hardware) {
	var p pixarray.Pixel
	p.R, p.G, p.B = 0, 0, 0
	for ledIdx := hardware.pa.NumPixels() - 1; ledIdx >= 0; ledIdx-- {
		hardware.pa.SetOne(ledIdx, p)
		hardware.pa.Write()
		time.Sleep(time.Millisecond * delayMs)
	}
}

func (b blink) switchOn(hardware *Hardware) {
	//TODO implement me
	panic("implement me")
}

func (b blink) switchOff(hardware *Hardware) {
	var p pixarray.Pixel
	p.R, p.G, p.B = 0, 0, 0
	hardware.pa.SetAll(p)
	hardware.pa.Write()
}

func (o opposite) switchOn(hardware *Hardware) {
	//TODO implement me
	panic("implement me")
}

func (o opposite) switchOff(hardware *Hardware) {
	//TODO implement me
	panic("implement me")
}

func (c chaseTail) switchOn(hardware *Hardware) {
	//TODO implement me
	panic("implement me")
}

func (c chaseTail) switchOff(hardware *Hardware) {
	//TODO implement me
	panic("implement me")
}

func (h hug) switchOn(hardware *Hardware) {
	//TODO implement me
	panic("implement me")
}

func (h hug) switchOff(hardware *Hardware) {
	//TODO implement me
	panic("implement me")
}

func (e explosion) switchOn(hardware *Hardware) {
	//TODO implement me
	panic("implement me")
}

func (e explosion) switchOff(hardware *Hardware) {
	//TODO implement me
	panic("implement me")
}

func toRgb(colorAndBrightness string) (r int, g int, b int) {
	r, g, b = colorconverter.HEX2RGB(colorAndBrightness[:6])
	bright64, err := strconv.ParseInt(colorAndBrightness[6:], 16, 0)
	if err != nil {
		panic(err)
	}
	r, g, b = scaleColor(r, g, b, float64(bright64)/255.0)
	return
}

func scaleColor(r int, g int, b int, scale float64) (int, int, int) {
	scaler := func(origin int) int {
		scaledFloat := float64(origin) * scale
		scaled := int(math.Round(scaledFloat))
		return scaled
	}

	return scaler(r), scaler(g), scaler(b)
}
