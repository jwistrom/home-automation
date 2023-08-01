package main

//
//import (
//	"fmt"
//	ws281x "github.com/rpi-ws281x/rpi-ws281x-go"
//	"github.com/stianeikeland/go-rpio/v4"
//	"log"
//	"time"
//)
//
//const (
//	brightness = 128
//	ledCounts  = 40
//	gpioPin    = 18
//	freq       = 800000
//	sleepTime  = 200
//)
//
//type ws struct {
//	ws2811 *ws281x.WS2811
//}
//
//func (ws *ws) init() error {
//	err := ws.ws2811.Init()
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (ws *ws) close() {
//	ws.ws2811.Fini()
//}
//
//func (ws *ws) lightOn() {
//	for x := 0; x < ledCounts; x++ {
//		ws.ws2811.Leds(0)[x] = uint32(0x00200000)
//	}
//	err := ws.ws2811.Render()
//	if err != nil {
//		panic(err)
//	}
//}
//
//func (ws *ws) lightOff() {
//	for x := 0; x < ledCounts; x++ {
//		ws.ws2811.Leds(0)[x] = uint32(0x00101010)
//	}
//	err := ws.ws2811.Render()
//	if err != nil {
//		panic(err)
//	}
//}
//
//func rgbToColor(r int, g int, b int) uint32 {
//	return uint32(r)<<16 | uint32(g)<<8 | uint32(b)
//}
//
//func main2() {
//	err := rpio.Open()
//	if err != nil {
//		panic(fmt.Sprint("unable to open gpio", err.Error()))
//	}
//
//	defer rpio.Close()
//
//	p14 := rpio.Pin(14)
//	p14.Output()
//	p14.High()
//
//	pin := rpio.Pin(18)
//	pin.Output()
//
//	log.Println("Toggling 20x")
//	for y := 0; y < 20; y++ {
//		for x := 0; x < 24; x++ {
//			if x%2 == 0 {
//				high(pin)
//			} else {
//				low(pin)
//			}
//		}
//	}
//
//}
//
//func high(pin rpio.Pin) {
//	pin.Low()
//	time.Sleep(time.Nanosecond * 700)
//	pin.High()
//	time.Sleep(time.Nanosecond * 600)
//}
//
//func low(pin rpio.Pin) {
//	pin.Low()
//	time.Sleep(time.Nanosecond * 350)
//	pin.High()
//	time.Sleep(time.Nanosecond * 800)
//}
//
//func main() {
//	options := ws281x.DefaultOptions
//	options.Channels[0].Brightness = brightness
//	options.Channels[0].LedCount = ledCounts
//	options.Channels[0].GpioPin = gpioPin
//	options.Channels[0].StripeType = ws281x.WS2812Strip
//	options.Frequency = freq
//
//	ws2811, err := ws281x.MakeWS2811(&options)
//	if err != nil {
//		panic(err)
//	}
//
//	ws := ws{
//		ws2811: ws2811,
//	}
//
//	err = ws.init()
//	if err != nil {
//		panic(err)
//	}
//	defer ws.close()
//
//	log.Println("Using lib")
//
//	for x := 0; x < 10; x++ {
//		log.Println("light on")
//		ws.lightOn()
//		time.Sleep(time.Second)
//		log.Println("light off")
//		ws.lightOff()
//		time.Sleep(time.Second * 2)
//	}
//
//}
