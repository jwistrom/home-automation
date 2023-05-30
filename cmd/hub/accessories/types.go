package accessories

import (
	"fmt"
)

type Accessory interface {
	getId() string
}

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
