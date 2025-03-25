package main

import (
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	position rl.Vector2
	speed    float32
	health   int
	weapon   int
}

type Enemy struct {
	position rl.Vector2
	speed    float32
	health   int
}

type Projectile struct {
	position rl.Vector2
	velocity rl.Vector2
	radius   float32
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

// ##################################################################
// ENEMY
// ##################################################################

func NewSimpleEnemy() *Enemy {
	return &Enemy{
		position: rl.NewVector2(float32(rand.Intn(screenWidth)), float32(rand.Intn(screenHeight))),
		speed:    20,
		health:   20,
	}
}

func (e *Enemy) Update() {
	e.moveTo(player.position)
}

func (e *Enemy) Draw() {
	rl.DrawCircleV(e.position, 15, enemyCol)
}

func (e *Enemy) moveTo(target rl.Vector2) {
	direction := rl.Vector2Subtract(target, e.position)
	distance := rl.Vector2Length(direction)
	direction = rl.Vector2Scale(direction, e.speed*rl.GetFrameTime()/distance)
	e.position = rl.Vector2Add(e.position, direction)
}

func (e *Enemy) IsCollidingWithProjectile(p *Projectile) bool {
	return rl.CheckCollisionCircles(e.position, 15, p.position, p.radius)
}

// ##################################################################
// PROJECTILE
// ##################################################################
func (p *Projectile) Update() {
	p.position = rl.Vector2Add(p.position, p.velocity)
}

func (p *Projectile) Draw() {
	rl.DrawCircleV(p.position, p.radius, rl.RayWhite)
}
