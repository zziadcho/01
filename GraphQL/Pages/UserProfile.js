import { Constructor } from "../Utilities/Constructor.js"
import { FetchQL } from "../Utilities/FetchData.js"
import { userInfo } from "../Utilities/QueriesQL.js"
export const UserProfile = async () => {
    const source = document.getElementById("source")
    const token = localStorage.getItem("JWT")
    const response = await FetchQL(userInfo, token)
    source.innerHTML = ""

    // const firstname = response.data.user[0].firstName
    // const lastname = response.data.user[0].lastName
    // const container

    const grid = Constructor("div", { id: "grid" }, source)
    const topBar = Constructor("div", { id: "top-bar" }, grid)

    const backbutton = Constructor("button", { id: "back-button", class : "material-symbols-outlined", textContent : "chevron_left"}, topBar)
    Constructor("h3", { id: "title", textContent: "Ready to go" }, topBar)
    Constructor("button", { id: "placeholder" }, topBar)

    const innerTitle = Constructor("div", { id: "inner-title", "textContent": "You're ready to go!" }, grid)
    Constructor("p", { "textContent": "You can now view your information:" }, innerTitle)

    const iconData = [
        { icon: "equalizer", text: "Your Level" },
        { icon: "reviews", text: "Audits Done" },
        { icon: "percent", text: "Audit Ratio" },
        { icon: "person", text: "Profile" },
        { icon: "calendar_month", text: "Schedule" },
        { icon: "handshake", text: "Collaboration" },
        { icon: "track_changes", text: "Progress" },
        { icon: "exposure_plus_2", text: "Boost" }
    ]

    iconData.forEach(item => {
        const iconGroup = Constructor("div", { class: "icon-group" }, grid)

        Constructor("span", {
            id: "icon",
            class: "material-symbols-outlined",
            textContent: item.icon
        }, iconGroup)

        Constructor("span", {
            class: "icon-text",
            textContent: item.text
        }, iconGroup)
    })

    const icons = document.querySelectorAll("icon")

    icons.addEventListener("click", function() {
        console.log("xd");
    })
}
