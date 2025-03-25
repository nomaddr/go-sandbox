package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	position rl.Vector2
	speed    float32
	health   int
	weapon   int
}

func NewPlayer(speed float32) *Player {
	return &Player{
		position: rl.NewVector2(screenWidth/2, screenHeight/2),
		speed:    speed,
		health:   100,
	}
}

func (p *Player) Update() {
	if rl.IsKeyDown(rl.KeyD) {
		p.position.X += p.speed * rl.GetFrameTime()
	}
	if rl.IsKeyDown(rl.KeyA) {
		p.position.X -= p.speed * rl.GetFrameTime()
	}
	if rl.IsKeyDown(rl.KeyW) {
		p.position.Y -= p.speed * rl.GetFrameTime()
	}
	if rl.IsKeyDown(rl.KeyS) {
		p.position.Y += p.speed * rl.GetFrameTime()
	}

	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		p.Shoot()
	}
}

func (p *Player) Draw() {
	rl.DrawCircleV(p.position, 20, playerCol)
}

func (p *Player) Shoot() {
	mousePos := rl.GetMousePosition()
	direction := rl.Vector2Subtract(mousePos, p.position)
	distance := float32(math.Sqrt(float64(direction.X*direction.X + direction.Y*direction.Y)))

	weaponSpeed := float32(100)
	if p.weapon == 0 {
		weaponSpeed = 1 // No weapon
	}

	direction = rl.Vector2Scale(direction, weaponSpeed/distance)

	projectiles = append(projectiles, &Projectile{
		position: p.position,
		velocity: direction,
		radius:   3,
	})
}
