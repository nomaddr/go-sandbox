package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Enemy struct {
	position rl.Vector2
	speed    float32
	health   int
}

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
