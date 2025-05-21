import { Constructor } from "../Utilities/Constructor.js"
import { PopUp } from "../Utilities/Popups.js"
import { UserProfile } from "./UserProfile.js"

export const Login = () => {
    const form = Constructor("form", { "id": "login-form" }, document.getElementById("source"))
    const usernameInput = Constructor("input", { "id": "username-input", "placeholder": "Username", "type": "text", "name": "username" }, form)
    const passwordInput = Constructor("input", { "id": "password-input", "placeholder": "Password", "type": "password", "name": "password" }, form)
    const button = Constructor("button", {"class": "material-symbols-outlined", "textContent": "chevron_right" }, form)

    form.addEventListener("submit", async (event) => {
        event.preventDefault()
        const username = document.getElementById("username-input").value,
            password = document.getElementById("password-input").value,
            credentials = btoa(`${username}:${password}`)

        const response = await fetch("https://learn.zone01oujda.ma/api/auth/signin", {
            method: "POST",
            headers: {
                "Authorization": `Basic ${credentials}`
            }
        })

        if (response.status != 200) {
            PopUp(response.status, "Invalid Credentials")
            return
        }
        const token = await response.json()

        localStorage.setItem("JWT", token)
        UserProfile()
    })
}