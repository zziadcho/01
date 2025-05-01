export const Router = () => {
    const path = window.location.pathname
    let view
    console.log(path);

    if (path === "/friends") {
        view = Friends()
    } else {
        view = Home()
    }

    document.getElementById("app").innerHTML = view
}

export const ElementConstructor = (tagName, id, className, textContent, appendTarget) => {
    const element = Object.assign(document.createElement(tagName), {
        id: id || "",
        className: className || "",
        textContent: textContent || ""
    })

    if (appendTarget && appendTarget.appendChild) appendTarget.appendChild(element)

    return element
}

const Home = () => {
    return "<h1>Home Page</h1>"
}

const Friends = () => {
    return "<h1>Friends List</h1>"
}
