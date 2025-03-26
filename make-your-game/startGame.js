//player stat&info
let playerX = 0, playerY = 0
let playerWidth = 250, playerHeight = 250
let playerHP = 10
let playerSpeed = 10

//enemy stat&info
let enemyWidth = 75, enemyHeight = 75
let enemyHP = 3
let enemySpeed = 5
let enemyDamage = 1
let enemyBuff = 1

//utilities
let mouseX, mouseY
let kbX = 0, kbY = 0, kbDuration = 0 //kb = knockback
let isInvis = false //Invis = invisibility
let invisDuration = 1500
const enemies = []
const maxSpawn = 0
const spawnInterval = 200

//basic movement
const startGame = () => {
    const pressedKeys = {}
    const button = document.getElementById("start")
    const player = Object.assign(document.createElement("img"), {
        id: "player",
        src: "assets/weaponR2.png"
    })
    button.style.display = "none"
    document.body.appendChild(player)

    //event listeners
    document.addEventListener("keydown", function (e) {
        pressedKeys[e.key] = true
    })
    document.addEventListener("keyup", function (e) {
        pressedKeys[e.key] = false
    })
    document.addEventListener("mousemove", function (e) {
        mouseX = e.clientX
        mouseY = e.clientY
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
    //move player 
    if (pressedKeys["w"]) playerY -= playerSpeed
    if (pressedKeys["a"]) playerX -= playerSpeed
    if (pressedKeys["s"]) playerY += playerSpeed
    if (pressedKeys["d"]) playerX += playerSpeed

    //stop player from going out of bound
    if (playerX < 0) playerX = 0
    if (playerY < 0) playerY = 0
    if (playerX > window.innerWidth - playerWidth) playerX = window.innerWidth - playerWidth
    if (playerY > window.innerHeight - playerHeight) playerY = window.innerHeight - playerHeight

    //player aim angle
    const MPDX = mouseX - (playerX + playerWidth / 2) //MPD = mousePlayerDistance
    const MPDY = mouseY - (playerY + playerHeight / 2)
    const aimAngle = Math.atan2(MPDY, MPDX) * (180 / Math.PI)
        
    //knockback player
    if (kbDuration > 0) {
        playerX += kbX
        playerY += kbY
        kbDuration--
        if (kbDuration <= 0) {
            kbX = 0
            kbY = 0
        }
    }

    //enemies
    enemies.forEach((enemy) => {
        const EPDX = playerX - enemy.x //EPD = enemyPlayerDistance
        const EPDY = playerY - enemy.y

        const EPD = Math.sqrt(EPDX * EPDX + EPDY * EPDY)

        if (EPD > 0) {
            enemy.x += (EPDX / EPD) * enemySpeed
            enemy.y += (EPDY / EPD) * enemySpeed
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

            const absX = Math.abs(EPDX)
            const absY = Math.abs(EPDY)

            let kbDistance = 150
            let kbFrames = 10

            if (absX > absY) {
                kbX = (EPDX < 0 ? -1 : 1) * (kbDistance / kbFrames)
                kbY = 0
            } else {
                kbX = 0
                kbY = (EPDY < 0 ? -1 : 1) * (kbDistance / kbFrames)
            }
            kbDuration = kbFrames
        }

        //enemy updates
        enemy.element.style.left = `${enemy.x}px`
        enemy.element.style.top = `${enemy.y}px`

    })
    //player updates
    player.style.left = `${playerX}px`
    player.style.top = `${playerY}px`
    player.style.transform = `rotate(${aimAngle}deg)`

    requestAnimationFrame(gameLoop)
}
requestAnimationFrame(gameLoop)
}