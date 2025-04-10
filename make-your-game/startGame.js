//player stat&info
const playerObj = {
    X: 700, Y: 400,
    Width: 75, Height: 175,
    HP: 10, Speed: 10, Money: 0, Weapon: null,
    kbX: 0, kbY: 0, kbDuration: 0
}

//enemy stat&info
const enemyObj = {
    // X: 0, Y: 0,
    Width: 75, Height: 175,
    bossHP: 10, HP: 3, Speed: 0.5,
    Damage: 1, Buff: 1.75
}

//bullet stat&info
const bulletObj = {
    X: playerObj.X + playerObj.Width / 2, Y: playerObj.Y + playerObj.Height / 2,
    Width: 15, Height: 15,
    Damage: 3, Speed: 40,
    fireRate: 10, fireRateInterval: null
}

//utilities
const bullets = [], enemies = [],
    mouseObj = {
        X: 0, Y: 0,
        buttonHeld: false,
    }

let invisDuration = 250,
    score = 0, isInvis = false, //Invis = invisibility
    roundStat = {
        roundCount: 0,
        normalRound: false,
        inLobby: true,
        maxSpawn: 2
    }

//actions
const spawnEnemies = () => {
    const DTS = 700 // DTS = distance to spawn
    if (roundStat.roundCount % 3 === 0) {
        let validPos = false,
            enemyX, enemyY
        while (!validPos) {
            enemyX = Math.random() * (window.innerWidth - enemyObj.Width),
                enemyY = Math.random() * (window.innerHeight - enemyObj.Height)

            const dx = enemyX - playerObj.X,
                dy = enemyY - playerObj.Y,
                distance = Math.sqrt(dx * dx + dy * dy)

            if (distance >= DTS) validPos = true
        }
        const enemyElement = Object.assign(document.createElement("img"), {
            id: "boss",
            src: "./assets/boss.png"
        })
        document.body.appendChild(enemyElement)
        enemies.push({
            element: enemyElement,
            x: enemyX,
            y: enemyY,
            type: "boss",
            hp: enemyObj.bossHP
        })
    }
    for (let i = 0; i < roundStat.maxSpawn; i++) {
        if (roundStat.normalRound) {
            let validPos = false,
                enemyX, enemyY
            while (!validPos) {
                enemyX = Math.random() * (window.innerWidth - enemyObj.Width),
                    enemyY = Math.random() * (window.innerHeight - enemyObj.Height)

                const dx = enemyX - playerObj.X,
                    dy = enemyY - playerObj.Y,
                    distance = Math.sqrt(dx * dx + dy * dy)

                if (distance >= DTS) validPos = true
            }

            const enemyElement = Object.assign(document.createElement("img"), {
                id: "enemy",
                src: "./assets/enemy.png"
            })
            document.body.appendChild(enemyElement)
            enemies.push({
                element: enemyElement,
                x: enemyX,
                y: enemyY,
                type: "troop",
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

const attack = (startX, startY, targetX, targetY) => {
    const dx = targetX - (startX + playerObj.Width / 2),
        dy = targetY - (startY + playerObj.Height / 2),
        bulletElement = Object.assign(document.createElement("div"), {
            id: "bullet"
        }),

        BPD = Math.sqrt(dx * dx + dy * dy), //BPD = bullet player distance
        dxNorm = dx / BPD,
        dyNorm = dy / BPD
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

const help = () => {
    alert(`- Move with WASD or Arrow keys
- Hold or Click the left mouse button to shoot bullets
- Move the player to the ready square to start the round
- Survive!`)
}

const startGame = () => {

    let startTime = 0
    const pressedKeys = {},
        runInfo = Object.assign(document.createElement("div"), {
            id: "runInfo"
        }),
        playerElement = Object.assign(document.createElement("img"), {
            id: "player",
            src: "./assets/player.png"
        }),
        shop = Object.assign(document.createElement("div"), {
            id: "shop",
            textContent: "Shop"
        }),
        ready = Object.assign(document.createElement("div"), {
            id: "ready",
            textContent: "Ready"
        }),
        menu = document.getElementById("menu")

    menu.remove()
    document.body.append(playerElement, runInfo, ready, shop)

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
        mouseObj.buttonHeld = true

        attack(playerObj.X, playerObj.Y, mouseObj.X, mouseObj.Y)

        bulletObj.fireRateInterval = setInterval(() => {
            if (mouseObj.buttonHeld) attack(playerObj.X, playerObj.Y, mouseObj.X, mouseObj.Y)
        }, bulletObj.fireRate);
    })
    document.addEventListener("mouseup", function (e) {
        mouseObj.buttonHeld = false
        if (bulletObj.fireRateInterval) {
            clearInterval(bulletObj.fireRateInterval)
            bulletObj.fireRateInterval = null
        }
    })
    document.addEventListener("dragstart", function (e) {
        e.preventDefault()
    })

    const gameLoop = () => {
        ////gameplay 
        //player
        const currentTime = performance.now(),
            timePassed = currentTime - startTime
        //move player 
        if (pressedKeys["w"] || pressedKeys["ArrowUp"]) playerObj.Y -= playerObj.Speed
        if (pressedKeys["a"] || pressedKeys["ArrowLeft"]) playerObj.X -= playerObj.Speed
        if (pressedKeys["s"] || pressedKeys["ArrowDown"]) playerObj.Y += playerObj.Speed
        if (pressedKeys["d"] || pressedKeys["ArrowRight"]) playerObj.X += playerObj.Speed

        //stop player from going out of bound
        if (playerObj.X < 0) playerObj.X = 0
        if (playerObj.Y < 0) playerObj.Y = 0
        if (playerObj.X > window.innerWidth - playerObj.Width) playerObj.X = window.innerWidth - playerObj.Width
        if (playerObj.Y > window.innerHeight - playerObj.Height) playerObj.Y = window.innerHeight - playerObj.Height

        //player aim angle
        // const MPDX = mouseObj.X - (playerObj.X + playerObj.Width / 2), //MPD = mousePlayerDistance
        // MPDY = mouseObj.Y - (playerObj.Y + playerObj.Height / 2),
        // aimAngle = Math.atan2(MPDY, MPDX) * (180 / Math.PI)

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

        const readyBounds = ready.getBoundingClientRect(),
            shopBounds = shop.getBoundingClientRect()

        //ready
        if (
            playerObj.X + playerObj.Width >= readyBounds.left &&
            playerObj.X <= readyBounds.left + readyBounds.width &&
            playerObj.Y + playerObj.Height >= readyBounds.top &&
            playerObj.Y <= readyBounds.top + readyBounds.height &&
            ready.parentNode
        ) {
            ready.remove(),
                shop.remove()
            startTime = performance.now()
            roundStat.normalRound = true
            roundStat.inLobby = false
            roundStat.roundCount++
            roundStat.maxSpawn *= enemyObj.Buff
            spawnEnemies()
        }

        if (
            playerObj.X + playerObj.Width >= shopBounds.left &&
            playerObj.X <= shopBounds.left + shopBounds.width &&
            playerObj.Y + playerObj.Height >= shopBounds.top &&
            playerObj.Y <= shopBounds.top + shopBounds.height &&
            shop.parentNode
        ) {
            document.getElementById("mask").style.display = "unset"
        } else {
            document.getElementById("mask").style.display = "none"
        }
        if (enemies.length === 0) roundStat.inLobby = true, roundStat.normalRound = false, roundStat.bossRound = false
        if (roundStat.inLobby) document.body.append(ready, shop)

        //enemies
        enemies.forEach((enemy, enemyIndex) => {

            const EPDX = playerObj.X - enemy.x, //EPD = enemyPlayerDistance
                EPDY = playerObj.Y - enemy.y,

                EPD = Math.sqrt(EPDX * EPDX + EPDY * EPDY)

            if (EPD > 0) {
                enemy.x += (EPDX / EPD) * enemyObj.Speed
                enemy.y += (EPDY / EPD) * enemyObj.Speed
            }

            //enemy/enemy collision check
            for (let i = 0; i < enemies.length; i++) {
                if (i != enemyIndex) {
                    const otherEnemy = enemies[i]
                    if (
                        enemy.x + enemyObj.Width >= otherEnemy.x &&
                        enemy.x <= otherEnemy.x + enemyObj.Width &&
                        enemy.y + enemyObj.Height >= otherEnemy.y &&
                        enemy.y <= otherEnemy.y + enemyObj.Height
                    ) {
                        const repelX = enemy.x - otherEnemy.x,
                            repelY = enemy.y - otherEnemy.y,
                            repelDist = Math.sqrt(repelX * repelX + repelY * repelY) || 1,
                            NRX = repelX / repelDist, //NR = normalized repel
                            NRY = repelY / repelDist,
                            repelStrength = 1.0

                        enemy.x += NRX * repelStrength
                        enemy.y += NRY * repelStrength

                        otherEnemy.x -= NRX * repelStrength
                        otherEnemy.y -= NRY * repelStrength
                    }
                }
            }

            //player/enemy collision check
            if (
                playerObj.X + playerObj.Width >= enemy.x &&
                playerObj.X <= enemy.x + enemyObj.Width &&
                playerObj.Y + playerObj.Height >= enemy.y &&
                playerObj.Y <= enemy.y + enemyObj.Height
            ) {
                damagePlayer(enemyObj.Damage)

                const absX = Math.abs(EPDX),
                    absY = Math.abs(EPDY)

                let kbDistance = 150,
                    kbFrames = 10

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
                    bullet.x + bulletObj.Width >= enemy.x &&
                    bullet.x <= enemy.x + enemyObj.Width &&
                    bullet.y + bulletObj.Height >= enemy.y &&
                    bullet.y <= enemy.y + enemyObj.Height
                ) {
                    enemy.hp -= bulletObj.Damage

                    if (bullet.element.parentNode) {
                        document.body.removeChild(bullet.element)
                        bullets.splice(bulletIndex, 1)
                    }

                    if (enemy.hp <= 0 && enemy.element.parentNode) {
                        score += 100
                        playerObj.Money += (enemy.type === "boss") ? 250 : (enemy.type === "troop") ? 25 : 0;
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
        // playerElement.style.transform = `rotate(${aimAngle}deg)`

        //run updates
        let runTime = `${(timePassed / 1000).toFixed(0)}s`
        runInfo.innerText = `
            Time: ${runTime}
            Round: ${roundStat.roundCount}
            HP: ${playerObj.HP}
            Money: ${playerObj.Money}
            Score: ${score}
        `

        if (playerObj.HP === 0) {
            alert(`You survived ${runTime}`);
        }

        requestAnimationFrame(gameLoop)
    }
    requestAnimationFrame(gameLoop)
}