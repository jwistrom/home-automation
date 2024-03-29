const axios = window.axios

function onTvBackgroundStateSwitch(evt) {
    console.log("Switch tv background light")
    axios.put("/tvbackgroundlight/state")
}

function onTvBackgroundModeChange(mode) {
    axios.put("tvbackgroundlight/mode", {}, {
        params: {
            mode: mode
        }
    })
}

function onTvBackgroundSpeedChange(speed) {
    axios.put("tvbackgroundlight/speed", {}, {
        params: {
            speed: speed
        }
    })
}

function onTvBackgroundColorChange(color) {
    axios.put("tvbackgroundlight/color", {}, {
        params: {
            color: color.substring(1)
        }
    })
}