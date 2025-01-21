const canvas = document.getElementById("main-screen")
const ctx = canvas.getContext("2d")
const clockTick = new Event('clockTick')

let up = false
let down = false
let left = false
let right = false
let space = false

document.addEventListener('keydown', function (e) {
    if (e.key == "w") {
        up = true
    }
    if (e.key == "s") {
        down = true
    }
    if (e.key == "a") {
        left = true
    }
    if (e.key == "d") {
        right = true
    }
    if (e.key == " ") {
        space = true
    }
})

document.addEventListener('keyup', function (e) {
    if (e.key == "w") {
        up = false
    }
    if (e.key == "s") {
        down = false
    }
    if (e.key == "a") {
        left = false
    }
    if (e.key == "d") {
        right = false
    }
    if (e.key == " ") {
        space = false
    }
})

function fireClock() {
    document.dispatchEvent(clockTick)
}
setInterval(fireClock, 16.99)


