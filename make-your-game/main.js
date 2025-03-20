//player stat&info
let playerX = 0, playerY = 0
let playerWidth = 50, playerHeight = 50 
const playerSpeed = 5

//basic movement
const startGame = () => {
    const player = Object.assign(document.createElement("div"), {
        id: "player",
    })
    const button = document.getElementById("start")
    button.style.display = "none"
    document.body.appendChild(player)

    const pressedKeys = {}

    document.addEventListener("keydown", function (e) {
        pressedKeys[e.key] = true
    })

    document.addEventListener("keyup", function (e) {
        pressedKeys[e.key] = false
    })

    const gameLoop = () => {
        if (pressedKeys['w']) {
            playerY -= playerSpeed
        }
        if (pressedKeys['d']) {
            playerX += playerSpeed
        }
        if (pressedKeys['s']) {
            playerY += playerSpeed
        }
        if (pressedKeys['a']) {
            playerX -= playerSpeed
        }

        if (playerX < 0) {
            playerX = 0
        }

        if (playerX > window.innerWidth - playerWidth) {
            playerX = window.innerWidth - playerWidth
        }

        if (playerY < 0) {
            playerY = 0
        }

        if (playerY > window.innerHeight - playerHeight) {
            playerY = window.innerHeight - playerHeight
        }

        console.log((player.style.width));
        

        player.style.left = `${playerX}px`
        player.style.top = `${playerY}px`

        requestAnimationFrame(gameLoop)
    }
    requestAnimationFrame(gameLoop)
}