package types

import "fmt"

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
	Mode  int
	Speed int
	Color string
}
