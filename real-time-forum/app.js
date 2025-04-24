import * as Functions from './functions.js'

document.querySelectorAll(".route").forEach(link => {
    link.addEventListener("click", function(e) {
        e.preventDefault()
        history.pushState(null, "", this.href)
        Functions.Router()
    })
});