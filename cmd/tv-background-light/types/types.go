package types

import "fmt"

type TvBackgroundLightMode int

const (
	Standard TvBackgroundLightMode = iota
	Circle
	Blink
	Opposite
	ChaseTail
	Hug
	Explosion
)

type SwitchState bool

func (ss SwitchState) Format(f fmt.State, c rune) {
	var text string
	if ss == On {
		text = "On"
	} else {
		text = "off"
	}
	_, _ = f.Write([]byte(text))
}

const (
	On  SwitchState = true
	Off SwitchState = false
)

type Status struct {
	State SwitchState
	Mode  TvBackgroundLightMode
	Speed int
	Color string
}
