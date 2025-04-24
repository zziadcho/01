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

const Home = () => {
    return "<h1>Home Page</h1>"
}

const Friends = () => {
    return "<h1>Friends List</h1>"
}