package accessories

type TvBackgroundLightMode int

const (
	Standard TvBackgroundLightMode = iota
	Circle
	Blink
	Opposite
	ChaseTail
	Hug
)

type TvBackgroundLightStatus struct {
	State SwitchState
	Mode  TvBackgroundLightMode
}

type TvBackgroundLight struct {
	state SwitchState
	mode  TvBackgroundLightMode
}

func (light *TvBackgroundLight) GetId() string {
	return "tv_background_light"
}

func (light *TvBackgroundLight) GetStatus() TvBackgroundLightStatus {
	return TvBackgroundLightStatus{
		State: light.state,
		Mode:  light.mode,
	}
}

func (light *TvBackgroundLight) SwitchState() SwitchState {
	if light.state == On {
		light.state = Off
	} else {
		light.state = On
	}

	return light.state
}

func (light *TvBackgroundLight) SetMode(mode TvBackgroundLightMode) {
	light.mode = mode
}
