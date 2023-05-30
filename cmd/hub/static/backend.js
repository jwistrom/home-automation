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