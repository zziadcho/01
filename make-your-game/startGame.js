//player stat&info
let playerX = 0, playerY = 0
let playerWidth = 50, playerHeight = 50
let playerHP = 10
const playerSpeed = 10

//enemy stat&info
let enemyWidth = 75, enemyHeight = 75
let enemyHP = 3
const enemySpeed = 5
const spawnInterval = 2000
const maxSpawn = 10
const enemies = []

//basic movement
const startGame = () => {
    const pressedKeys = {}
    const button = document.getElementById("start")
    const player = Object.assign(document.createElement("div"), {
        id: "player",
    })
    button.style.display = "none"
    document.body.appendChild(player)

    document.addEventListener("keydown", function (e) {
        pressedKeys[e.key] = true
    })
    document.addEventListener("keyup", function (e) {
        pressedKeys[e.key] = false
    })
    setInterval(() => {
        if (enemies.length < maxSpawn) {
            let enemyX = Math.random() * (window.innerWidth - enemyWidth)
            let enemyY = Math.random() * (window.innerHeight - enemyHeight)
            const enemy = Object.assign(document.createElement("div"), {
                id: "enemy"
            })
            document.body.appendChild(enemy)
            enemies.push({
                element: enemy,
                x: enemyX,
                y: enemyY,
                hp: enemyHP
            })
        }
    }, spawnInterval)

    const gameLoop = () => {

        if (pressedKeys["w"]) playerY -= playerSpeed
        if (pressedKeys["d"]) playerX += playerSpeed
        if (pressedKeys["s"]) playerY += playerSpeed
        if (pressedKeys["a"]) playerX -= playerSpeed

        if (playerX < 0) playerX = 0
        if (playerY < 0) playerY = 0
        if (playerX > window.innerWidth - playerWidth) playerX = window.innerWidth - playerWidth
        if (playerY > window.innerHeight - playerHeight) playerY = window.innerHeight - playerHeight

        enemies.forEach((enemy) => {
            const xDistance = playerX - enemy.x
            const yDistance = playerY - enemy.y
            const distance = Math.sqrt(xDistance * xDistance + yDistance * yDistance)

            if (distance > 0) {
                enemy.x += (xDistance / distance) * enemySpeed
                enemy.y += (yDistance / distance) * enemySpeed
            }

            enemy.element.style.left = `${enemy.x}px`
            enemy.element.style.top = `${enemy.y}px`

            if (playerX < enemy.x + enemyWidth &&
                playerX + playerWidth > enemy.x &&
                playerY < enemy.y + enemyHeight &&
                playerY + playerHeight > enemy.y) {
                console.log("collided")
            }
        });

        player.style.left = `${playerX}px`
        player.style.top = `${playerY}px`

        requestAnimationFrame(gameLoop)
    }
    requestAnimationFrame(gameLoop)
}