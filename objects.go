package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Projectile struct {
	position rl.Vector2
	velocity rl.Vector2
	radius   float32
}

func (p *Projectile) Update() {
	p.position = rl.Vector2Add(p.position, p.velocity)
}

func (p *Projectile) Draw() {
	rl.DrawCircleV(p.position, p.radius, rl.RayWhite)
}
