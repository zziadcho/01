//player stat&info
let playerX = 0, playerY = 0,
    playerWidth = 50, playerHeight = 50,
    playerHP = 10, playerSpeed = 10

//enemy stat&info
let enemyWidth = 75, enemyHeight = 75,
    enemyHP = 3, enemySpeed = 0.1
enemyDamage = 1, enemyBuff = 1

//bullet stat&info
let bulletX = playerX + playerWidth / 2, bulletY = playerY + playerHeight / 2,
    bulletWidth = 15, bulletHeight = 15
bulletDamage = 1, bulletSpeed = 10

//utilities
const bullets = [], enemies = []
const maxSpawn = 1, spawnInterval = 200
let mouseX, mouseY,
    kbX = 0, kbY = 0, kbDuration = 0, //kb = knockback
    invisDuration = 1500, isInvis = false //Invis = invisibility

const startGame = () => {
    const pressedKeys = {}
    const button = document.getElementById("start")
    const player = Object.assign(document.createElement("img"), {
        id: "player",
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
    document.addEventListener("click", function (e) {
        shootBullet(playerX, playerY, mouseX, mouseY)
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

    const shootBullet = (startX, startY, targetX, targetY) => {
        const dx = targetX - (startX + playerWidth / 2)
        const dy = targetY - (startY + playerHeight / 2)
        const bullet = Object.assign(document.createElement("div"), {
            id: "bullet"
        })

        const BPD = Math.sqrt(dx * dx + dy * dy) //BPD = bullet player
        const dxNorm = dx / BPD
        const dyNorm = dy / BPD
        document.body.appendChild(bullet)
        bullets.push({
            element: bullet,
            x: startX + playerWidth / 2,
            y: startY + playerHeight / 2,
            dx: dxNorm * bulletSpeed,
            dy: dyNorm * bulletSpeed,
            damage: bulletDamage
        })
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
        enemies.forEach((enemy, enemyIndex) => {
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

            bullets.forEach((bullet, bulletIndex) => {
                if (bullet.x + bulletWidth >= enemy.x
                    && bullet.x <= enemy.x + enemyWidth
                    && bullet.y + bulletHeight >= enemy.y
                    && bullet.y <= enemy.y + enemyHeight
                ) {
                    // Reduce enemy HP
                    enemy.hp -= bullet.damage
                    console.log(enemy.hp);
                    

                    // Remove bullet
                    document.body.removeChild(bullet.element)
                    bullets.splice(bulletIndex, 1)

                    // Remove enemy if HP reaches 0
                    if (enemy.hp <= 0) {
                        document.body.removeChild(enemy.element)
                        enemies.splice(enemyIndex, 1)
                    }
                }
            });

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
        });
        //player updates
        player.style.left = `${playerX}px`
        player.style.top = `${playerY}px`
        player.style.transform = `rotate(${aimAngle}deg)`

        if (playerHP === 0) {
            alert("gameOver")
        }

        requestAnimationFrame(gameLoop)
    }
    requestAnimationFrame(gameLoop)
}