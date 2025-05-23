export const Constructor = (tag, attribues = {}, appendTarget) => {
    try {
        const element = document.createElement(tag)
        for (const [key, value] of Object.entries(attribues)) {
            if (key === "textContent") element.textContent = value
            element.setAttribute(key, value)
        }
        if (appendTarget && appendTarget.appendChild) {
            if (appendTarget === document) { document.body.appendChild(element) } else { appendTarget.appendChild(element) }
        } return element
    }
    catch (err) {
        console.log(err.message)
    }
}