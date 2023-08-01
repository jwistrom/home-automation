package main

import (
	cc "github.com/SimonWaldherr/ColorConverterGo"
	"github.com/SimonWaldherr/ws2812/pixarray"
	"log"
	"time"
)

const (
	ledCount   = 100
	colorCount = 3
	ledPin     = 18
	defaultDma = 10
	frequency  = 800000
)

func main() {

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
	time.Sleep(1500 * time.Millisecond)
	//________________________________________---
	lcolours := []int{1, 2, 3}
	ledarray := [ledCount]int{}
	for {
		for colorpicker := 0; colorpicker < 3; colorpicker++ {
			for ledIdx := 0; ledIdx < ledCount; ledIdx++ {
				ledarray[ledIdx] = lcolours[colorpicker]
			}
		}

	}

	//____________________________________________
	colorPos := []int{-1, -2, -3}

	for {
		for ledIdx := 0; ledIdx < ledCount; ledIdx++ {
			if colorPos[0] == ledIdx {
				p.R, p.G, p.B = cc.HEX2RGB("ff0000")
			} else if colorPos[1] == ledIdx {
				p.R, p.G, p.B = cc.HEX2RGB("00ff00")
			} else if colorPos[2] == ledIdx {
				p.R, p.G, p.B = cc.HEX2RGB("0000ff")
			} else {
				p.R, p.G, p.B = cc.HEX2RGB("000000")
			}
			pa.SetOne(ledIdx, p)
		}
		pa.Write()

		for c := 0; c < len(colorPos); c++ {
			if colorPos[c] > ledCount-1 {
				colorPos[c] = 0
			} else {
				colorPos[c] = colorPos[c] + 1
			}
		}
		time.Sleep(50 * time.Millisecond)
	}

	//for {
	//	for j := 0; j < 360; j++ {
	//		for i := 0; i < pa.NumPixels(); i++ {
	//			h := int(360.0 / float32(pa.NumPixels()) * float32(i))
	//			p.R, p.G, p.B = cc.HSV2RGB(h+j, 100, 100)
	//			pa.SetOne((i)%pa.NumPixels(), p)
	//		}
	//		time.Sleep(2 * time.Millisecond)
	//		pa.Write()
	//	}
	//}
}
