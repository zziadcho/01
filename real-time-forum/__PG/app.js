import * as Functions from './functions.js'

const container = document.getElementById("main-container");
const profileSelect = Functions.ElementConstructor("div", "profile-select", "", "", container); //parent div

const profileImg = Functions.ElementConstructor("img", "profile-pic", "", "", profileSelect); //1st element
const selectContainer = Functions.ElementConstructor("div", "select-container", "", "", profileSelect); //2nd element - another div

const selectElement = Functions.ElementConstructor("select", "user-select", "", "", selectContainer); //1st child
const option1 = Functions.ElementConstructor("option", "", "", "zziadcho", selectElement);
const option2 = Functions.ElementConstructor("option", "", "", "Sign Out", selectElement);

const dropdownArrow = Functions.ElementConstructor("span", "", "dropdown-arrow", "▼", selectContainer);//2nd child

document.querySelectorAll(".route").forEach(link => {
    link.addEventListener("click", function (e) {
        e.preventDefault()
        history.pushState(null, "", this.href)
        Functions.Router()
    })
})

