package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 800
	screenHeight = 450
)

var (
	player      *Player
	enemies     []*Enemy
	projectiles []*Projectile

	palette   = rl.LoadImage("res/palettes/rust-gold-8-1x.png")
	colors    = rl.LoadImageColors(palette)
	playerCol = colors[4]
	enemyCol  = rl.Red
	bgColor   = colors[7]
)

func Init() {
	rl.InitWindow(screenWidth, screenHeight, "Top-Down Shooter")
	rl.SetTargetFPS(60)

	player = NewPlayer(200)

	// Spawn initial enemies
	for i := 0; i < 5; i++ {
		enemies = append(enemies, NewSimpleEnemy())
	}
}

func Update() {
	player.Update()

	for _, enemy := range enemies {
		enemy.Update()
	}

	for _, projectile := range projectiles {
		projectile.Update()
	}

	checkCollisions()
}

func Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(bgColor)

	player.Draw()

	for _, enemy := range enemies {
		enemy.Draw()
	}

	for _, projectile := range projectiles {
		projectile.Draw()
	}

	rl.DrawText(fmt.Sprintf("Health: %d", player.health), 10, 10, 20, rl.RayWhite)

	rl.EndDrawing()
}

func checkCollisions() {
	// Check player-enemy collisions
	for i := range enemies {
		if rl.CheckCollisionCircles(player.position, 20, enemies[i].position, 15) {
			player.health -= 10
			if player.health <= 0 {
				player.health = 0
				// Handle player death (e.g., reset game, show game over screen)
			}
		}
	}

	// Check enemy-enemy collision
	for i := 1; i < len(enemies); i++ {
		if rl.CheckCollisionCircles(enemies[i].position, 15, enemies[i-1].position, 15) {

		}

	}

	// Check projectile-enemy collisions
	for i := 0; i < len(projectiles); i++ {
		for j := 0; j < len(enemies); j++ {
			if enemies[j].IsCollidingWithProjectile(projectiles[i]) {
				// Remove projectile
				projectiles = append(projectiles[:i], projectiles[i+1:]...)
				i--

				// Damage enemy
				enemies[j].health -= 10
				if enemies[j].health <= 0 {
					// Remove enemy
					enemies = append(enemies[:j], enemies[j+1:]...)
					j--
				}
				break
			}
		}
	}

	// Remove projectiles that go off-screen
	for i := 0; i < len(projectiles); i++ {
		if projectiles[i].position.X < 0 || projectiles[i].position.X > screenWidth ||
			projectiles[i].position.Y < 0 || projectiles[i].position.Y > screenHeight {
			projectiles = append(projectiles[:i], projectiles[i+1:]...)
			i--
		}
	}
}
