//player stat&info
const playerObj = {
    X: 0, Y: 0,
    Width: 50, Height: 50,
    HP: 10, Speed: 10,
    kbX: 0, kbY: 0, kbDuration: 0
}

//enemy stat&info
const enemyObj = {
    // X: 0, Y: 0,
    Width: 75, Height: 75,
    HP: 3, Speed: 2,
    Damage: 1, Buff: 1
}

//bullet stat&info
const bulletObj = {
    X: playerObj.X + playerObj.Width / 2, Y: playerObj.Y + playerObj.Height / 2,
    Width: 15, Height: 15,
    Damage: 1, Speed: 10
}

//utilities
const bullets = [], enemies = []
const maxSpawn = 5, spawnInterval = 200
const mouseObj = {
    X: 0,
    Y: 0
}

let invisDuration = 1500,
    isInvis = false //Invis = invisibility

const startGame = () => {
    const pressedKeys = {}
    const button = document.getElementById("start")
    const playerElement = Object.assign(document.createElement("div"), {
        id: "player",
    })
    button.style.display = "none"
    document.body.appendChild(playerElement)

    //event listeners
    document.addEventListener("keydown", function (e) {
        pressedKeys[e.key] = true
    })
    document.addEventListener("keyup", function (e) {
        pressedKeys[e.key] = false
    })
    document.addEventListener("mousemove", function (e) {
        mouseObj.X = e.clientX
        mouseObj.Y = e.clientY
    })
    document.addEventListener("mousedown", function (e) {
        shootBullet(playerObj.X, playerObj.Y, mouseObj.X, mouseObj.Y)
    })
    setInterval(() => {
        if (enemies.length < maxSpawn) {
            let enemyX = Math.random() * (window.innerWidth - enemyObj.Width),
                enemyY = Math.random() * (window.innerHeight - enemyObj.Height)
            const enemyElement = Object.assign(document.createElement("div"), {
                id: "enemy"
            })
            document.body.appendChild(enemyElement)
            enemies.push({
                element: enemyElement,
                x: enemyX,
                y: enemyY,
                hp: enemyObj.HP
            })
        }
    }, spawnInterval)
    //actions
    const damagePlayer = (damage) => {
        if (isInvis) return
        playerObj.HP -= damage

        isInvis = true
        setTimeout(() => {
            isInvis = false
        }, invisDuration)
    }

    const shootBullet = (startX, startY, targetX, targetY) => {
        const dx = targetX - (startX + playerObj.Width / 2)
        const dy = targetY - (startY + playerObj.Height / 2)
        const bulletElement = Object.assign(document.createElement("div"), {
            id: "bullet"
        })

        const BPD = Math.sqrt(dx * dx + dy * dy) //BPD = bullet player
        const dxNorm = dx / BPD
        const dyNorm = dy / BPD
        document.body.appendChild(bulletElement)
        bullets.push({
            element: bulletElement,
            x: startX + playerObj.Width / 2,
            y: startY + playerObj.Height / 2,
            dx: dxNorm * bulletObj.Speed,
            dy: dyNorm * bulletObj.Speed,
            damage: bulletObj.Damage
        })
    }

    const gameLoop = () => {

        //move player 
        if (pressedKeys["w"]) playerObj.Y -= playerObj.Speed
        if (pressedKeys["a"]) playerObj.X -= playerObj.Speed
        if (pressedKeys["s"]) playerObj.Y += playerObj.Speed
        if (pressedKeys["d"]) playerObj.X += playerObj.Speed

        //stop player from going out of bound
        if (playerObj.X < 0) playerObj.X = 0
        if (playerObj.Y < 0) playerObj.Y = 0
        if (playerObj.X > window.innerWidth - playerObj.Width) playerObj.X = window.innerWidth - playerObj.Width
        if (playerObj.Y > window.innerHeight - playerObj.Height) playerObj.Y = window.innerHeight - playerObj.Height

        //player aim angle
        const MPDX = mouseObj.X - (playerObj.X + playerObj.Width / 2) //MPD = mousePlayerDistance
        const MPDY = mouseObj.Y - (playerObj.Y + playerObj.Height / 2)
        const aimAngle = Math.atan2(MPDY, MPDX) * (180 / Math.PI)

        //knockback player
        if (playerObj.kbDuration > 0) {
            playerObj.X += playerObj.kbX
            playerObj.Y += playerObj.kbY
            playerObj.kbDuration--
            if (playerObj.kbDuration <= 0) {
                playerObj.kbX = 0
                playerObj.kbY = 0
            }
        }

        //enemies
        enemies.forEach((enemy, enemyIndex) => {

            const EPDX = playerObj.X - enemy.x //EPD = enemyPlayerDistance
            const EPDY = playerObj.Y - enemy.y

            const EPD = Math.sqrt(EPDX * EPDX + EPDY * EPDY)

            if (EPD > 0) {
                enemy.x += (EPDX / EPD) * enemyObj.Speed
                enemy.y += (EPDY / EPD) * enemyObj.Speed
            }

            //enemy/enemy collision check
            for (let i = 0; i < enemies.length; i++) {
                if (i != enemyIndex) {
                    const otherEnemy = enemies[i];
                    if (enemy.x + enemyObj.Width >= otherEnemy.X
                        &&
                        enemy.x <= otherEnemy.X + enemyObj.Width
                        &&
                        enemy.y + enemyObj.Height >= otherEnemy.Y
                        &&
                        enemy.y <= otherEnemy.Y + enemyObj.Height
                    ) {
                        console.log('enemies collided');
                    }
                }
            }
            //player/enemy collision check
            if (playerObj.X + playerObj.Width >= enemy.x
                &&
                playerObj.X <= enemy.x + enemyObj.Width
                &&
                playerObj.Y + playerObj.Height >= enemy.y
                &&
                playerObj.Y <= enemy.y + enemyObj.Height
            ) {
                damagePlayer(enemyObj.Damage)

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

            //bullet/enemy collision check
            bullets.forEach((bullet, bulletIndex) => {
                if (bullet.x + bulletObj.Width >= enemy.x
                    && bullet.x <= enemy.x + enemyObj.Width
                    && bullet.y + bulletObj.Height >= enemy.y
                    && bullet.y <= enemy.y + enemyObj.Height
                ) {
                    enemy.hp -= bulletObj.Damage

                    if (bullet.element.parentNode) {
                        document.body.removeChild(bullet.element)
                        bullets.splice(bulletIndex, 1)
                    }

                    if (enemy.hp <= 0 && enemy.element.parentNode) {
                        document.body.removeChild(enemy.element)
                        enemies.splice(enemyIndex, 1)
                    }
                }
            })

            //enemy updates
            enemy.element.style.left = `${enemy.x}px`
            enemy.element.style.top = `${enemy.y}px`
        })

        bullets.forEach((bullet, index) => {
            bullet.x += bullet.dx
            bullet.y += bullet.dy

            if (bullet.x < 0 || bullet.x > window.innerWidth ||
                bullet.y < 0 || bullet.y > window.innerHeight) {
                document.body.removeChild(bullet.element)
                bullets.splice(index, 1)
                return
            }

            bullet.element.style.left = `${bullet.x}px`
            bullet.element.style.top = `${bullet.y}px`
        })
        //player updates
        playerElement.style.left = `${playerObj.X}px`
        playerElement.style.top = `${playerObj.Y}px`
        playerElement.style.transform = `rotate(${aimAngle}deg)`

        if (playerObj.HP === 0) {
            alert("gameOver")
        }

        requestAnimationFrame(gameLoop)
    }
    requestAnimationFrame(gameLoop)
}