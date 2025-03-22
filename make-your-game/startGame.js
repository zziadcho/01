//player stat&info
let playerX = 0, playerY = 0
let playerWidth = 50, playerHeight = 50
let playerHP = 10
let playerSpeed = 10

//enemy stat&info
let enemyWidth = 75, enemyHeight = 75
let enemyHP = 3
let enemySpeed = 5
let enemyDamage = 1
let enemyBuff = 1

//utilities
let isInvis = false //Invis = invisibility
let invisDuration = 1500
const spawnInterval = 200
const maxSpawn = 1
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

    //actions
    const damagePlayer = (damage) => {
        if (isInvis) return
        playerHP -= damage

        isInvis = true
        setTimeout(() => {
            isInvis = false
        }, invisDuration)
}
const gameLoop = () => {
    //keys pressed
    if (pressedKeys["w"]) playerY -= playerSpeed
    if (pressedKeys["a"]) playerX -= playerSpeed
    if (pressedKeys["s"]) playerY += playerSpeed
    if (pressedKeys["d"]) playerX += playerSpeed

    //arena bounderies
    if (playerX < 0) playerX = 0
    if (playerY < 0) playerY = 0
    if (playerX > window.innerWidth - playerWidth) playerX = window.innerWidth - playerWidth
    if (playerY > window.innerHeight - playerHeight) playerY = window.innerHeight - playerHeight

    //enemies
    enemies.forEach((enemy) => {
        const distanceX = playerX - enemy.x
        const distanceY = playerY - enemy.y

        const distance = Math.sqrt(distanceX * distanceX + distanceY * distanceY)

        if (distance > 0) {
            enemy.x += (distanceX / distance) * enemySpeed
            enemy.y += (distanceY / distance) * enemySpeed
        }

        //collision check
        if (playerX + playerWidth >= enemy.x
            &&
            playerX <= enemy.x + enemyWidth
            &&
            playerY + playerHeight >= enemy.y
            &&
            playerY <= enemy.y + enemyHeight
        ) {
            damagePlayer(enemyDamage)
            const absX = Math.abs(distanceX)
            const absY = Math.abs(distanceY)

            if (absX > absY) {
                if (distanceX < 0) {
                    playerX -= 75
                } else {
                    playerX += 75
                }
            } else {
                if (distanceY < 0) {
                    playerY -= 75
                } else {
                    playerY += 75
                }
            }
        }

        enemy.element.style.left = `${enemy.x}px`
        enemy.element.style.top = `${enemy.y}px`

    })
    //player
    player.style.left = `${playerX}px`
    player.style.top = `${playerY}px`

    requestAnimationFrame(gameLoop)
}
requestAnimationFrame(gameLoop)
}