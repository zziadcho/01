//player stat&info
const playerObj = {
    X: 500, Y: 500,
    Width: 50, Height: 50,
    HP: 10, Speed: 10,
    kbX: 0, kbY: 0, kbDuration: 0
}

//enemy stat&info
const enemyObj = {
    // X: 0, Y: 0,
    Width: 75, Height: 75,
    HP: 3, Speed: 0.5,
    Damage: 1, Buff: 1.75
}

//bullet stat&info
const bulletObj = {
    X: playerObj.X + playerObj.Width / 2, Y: playerObj.Y + playerObj.Height / 2,
    Width: 15, Height: 15,
    Damage: 3, Speed: 40
}

//utilities
const bullets = [], enemies = [],
    mouseObj = {
        X: 0, Y: 0
    }

let invisDuration = 250,
    score = 0, isInvis = false, //Invis = invisibility
    roundStat = {
        roundCount: 0,
        bossRound: false,
        normalRound: false,
        inLobby: true,
        maxSpawn: 2
    }

//actions
const spawnEnemies = () => {
    for (let i = 0; i < roundStat.maxSpawn; i++) {
        if (roundStat.normalRound) {
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
    }
}

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

const startGame = () => {

    let startTime = 0
    const pressedKeys = {},
        runInfo = Object.assign(document.createElement("div"), {
            id: "runInfo"
        }),
        playerElement = Object.assign(document.createElement("div"), {
            id: "player",
        }),
        shop = document.getElementById("shop"),
        ready = document.getElementById("ready"),
        menu = document.getElementById("menu")

    menu.remove()
    document.body.append(playerElement, runInfo, shop, ready)

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

    const gameLoop = () => {
        ////gameplay 
        //player
        const currentTime = performance.now()
        const timePassed = currentTime - startTime
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

        const readyBounds = ready.getBoundingClientRect()


        //ready
        if (
            playerObj.X + playerObj.Width >= readyBounds.left
            &&
            playerObj.X <= readyBounds.left + readyBounds.width
            &&
            playerObj.Y + playerObj.Height >= readyBounds.top
            &&
            playerObj.Y <= readyBounds.top + readyBounds.height
            &&
            ready.parentNode
        ) {
            ready.remove(), shop.remove()
            startTime = performance.now()
            roundStat.normalRound = true
            roundStat.inLobby = false
            roundStat.roundCount++
            roundStat.maxSpawn *= enemyObj.Buff
            spawnEnemies()
        }
        if (enemies.length === 0) roundStat.inLobby = true, roundStat.normalRound = false, roundStat.bossRound = false

        if (roundStat.inLobby && !roundStat.bossRound && !roundStat.normalRound) document.body.append(ready, shop)
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
                    const otherEnemy = enemies[i]
                    if (
                        enemy.x + enemyObj.Width >= otherEnemy.x
                        &&
                        enemy.x <= otherEnemy.x + enemyObj.Width
                        &&
                        enemy.y + enemyObj.Height >= otherEnemy.y
                        &&
                        enemy.y <= otherEnemy.y + enemyObj.Height
                    ) {
                        const repelX = enemy.x - otherEnemy.x
                        const repelY = enemy.y - otherEnemy.y

                        const repelDist = Math.sqrt(repelX * repelX + repelY * repelY) || 1
                        const NRX = repelX / repelDist //NR = normalized repel
                        const NRY = repelY / repelDist

                        const repelStrength = 1.0
                        enemy.x += NRX * repelStrength
                        enemy.y += NRY * repelStrength

                        otherEnemy.x -= NRX * repelStrength
                        otherEnemy.y -= NRY * repelStrength
                    }
                }
            }

            //player/enemy collision check
            if (
                playerObj.X + playerObj.Width >= enemy.x
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
                    playerObj.kbX = (EPDX < 0 ? -1 : 1) * (kbDistance / kbFrames)
                    playerObj.kbY = 0
                } else {
                    playerObj.kbX = 0
                    playerObj.kbY = (EPDY < 0 ? -1 : 1) * (kbDistance / kbFrames)
                }
                playerObj.kbDuration = kbFrames
            }

            //bullet/enemy collision check
            bullets.forEach((bullet, bulletIndex) => {
                if (
                    bullet.x + bulletObj.Width >= enemy.x
                    &&
                    bullet.x <= enemy.x + enemyObj.Width
                    &&
                    bullet.y + bulletObj.Height >= enemy.y
                    &&
                    bullet.y <= enemy.y + enemyObj.Height
                ) {
                    enemy.hp -= bulletObj.Damage

                    if (bullet.element.parentNode) {
                        document.body.removeChild(bullet.element)
                        bullets.splice(bulletIndex, 1)
                    }

                    if (enemy.hp <= 0 && enemy.element.parentNode) {
                        score += 100
                        document.body.removeChild(enemy.element)
                        enemies.splice(enemyIndex, 1)
                    }
                }
            })

            //enemy updates
            enemy.element.style.left = `${enemy.x}px`
            enemy.element.style.top = `${enemy.y}px`
        })

        //bullets
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

        //run updates
        let runTime = `${(timePassed / 1000).toFixed(0)}s`
        runInfo.innerText = `
            Time: ${runTime}
            Round: ${roundStat.roundCount}
            HP: ${playerObj.HP}
            Score: ${score}
        `

        if (playerObj.HP === 0) {
            alert(`You survived ${runTime}`);
        }

        ////interaction


        requestAnimationFrame(gameLoop)
    }
    requestAnimationFrame(gameLoop)
}