package main

import (
	"fmt"
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameObject interface {
	Update()
	Draw()
}

type Player struct {
	position                            rl.Vector2
	atlas                               rl.Texture2D
	src, dest                           rl.Rectangle
	speed                               float32
	health, weapon, curFrame, numFrames int
}

type Enemy struct {
	position rl.Vector2
	speed    float32
	health   int
}

type Projectile struct {
	position rl.Vector2
	velocity rl.Vector2
	speed    float32
	radius   float32
}

const (
	screenWidth  = 800
	screenHeight = 450
)

var (
	player      *Player
	enemies     []*Enemy
	projectiles []*Projectile
	red         = rl.NewColor(146, 33, 95, 100)
	purple      = rl.NewColor(89, 12, 104, 100)
	black       = rl.NewColor(0, 8, 40, 100)
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Top-Down Shooter")
	rl.SetTargetFPS(60)

	player = NewPlayer(200)

	// Spawn initial enemies
	for i := 0; i < 5; i++ {
		enemies = append(enemies, NewSimpleEnemy())
	}

	for !rl.WindowShouldClose() {
		update()
		draw()
	}

	rl.CloseWindow()
}

func update() {
	player.Update()

	for i := 0; i < len(enemies); i++ {
		enemies[i].Update()
	}

	for i := 0; i < len(projectiles); i++ {
		projectiles[i].Update()
	}

	checkCollisions()
}

func draw() {
	rl.BeginDrawing()
	rl.ClearBackground(black)

	player.Draw()

	for _, enemy := range enemies {
		enemy.Draw()
	}

	for _, projectile := range projectiles {
		projectile.Draw()
	}

	// Draw player health
	rl.DrawText(fmt.Sprintf("Health: %d", player.health), 10, 10, 20, rl.Black)

	rl.EndDrawing()
}

func NewSimpleEnemy() *Enemy {
	return &Enemy{
		position: rl.NewVector2(float32(rand.Intn(screenWidth)), float32(rand.Intn(screenHeight))),
		speed:    20,
		health:   20,
	}
}

func NewEnemy(position rl.Vector2, speed float32, health int) *Enemy {
	return &Enemy{
		position: position,
		speed:    speed,
		health:   health,
	}
}

func NewPlayer(speed float32) *Player {
	p := new(Player)
	p.position = rl.NewVector2(screenWidth/2, screenHeight/2)
	p.speed = speed
	p.health = 100
	p.atlas = rl.LoadTexture("res/idle2.png")
	p.numFrames = 6
	p.src = rl.NewRectangle(0, 0, float32(p.atlas.Width)/float32(p.numFrames), float32(p.atlas.Height))
	p.dest = rl.NewRectangle(p.position.X, p.position.Y, p.src.Width*2, p.src.Height*2)

	return p
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
	player.dest.X = player.position.X
	player.dest.Y = player.position.Y
}

func (p *Player) Draw() {
	rl.DrawCircleV(p.position, 20, purple)
	rl.DrawCircleLines(int32(p.position.X), int32(p.position.Y), 20, red)

	//rl.DrawTexturePro(player.atlas, player.src, player.dest, rl.Vector2{player.dest.Width / 2, player.dest.Height / 2}, 1, rl.White)
	//rl.DrawRectangleLines(int32(dest.X), int32(dest.Y), int32(dest.Width), int32(dest.Height), rl.Green)
}

func (p *Player) Shoot() {
	mousePos := rl.GetMousePosition()
	direction := rl.Vector2Subtract(mousePos, p.position)
	distance := float32(math.Sqrt(float64(direction.X*direction.X + direction.Y*direction.Y)))
	var weaponSpeed float32
	switch p.weapon {
	case 0:
		weaponSpeed = 1 // no weapon
	default:
		weaponSpeed = 100
	}
	direction = rl.Vector2Scale(direction, weaponSpeed/distance)

	projectile := &Projectile{
		position: p.position,
		velocity: direction,
		radius:   3,
	}
	projectiles = append(projectiles, projectile)
}

func (e *Enemy) moveTo(pos rl.Vector2) {
	direction := rl.Vector2Subtract(pos, e.position)
	distance := float32(math.Sqrt(float64(direction.X*direction.X + direction.Y*direction.Y)))
	direction = rl.Vector2Scale(direction, e.speed*rl.GetFrameTime()/distance)
	e.position = rl.Vector2Add(e.position, direction)
}

func (e *Enemy) Update() {
	e.moveTo(player.position)
}

func (e *Enemy) Draw() {
	rl.DrawCircleV(e.position, 15, red)
}

func (e *Enemy) IsCollidingWithProjectile(p *Projectile) bool {
	return rl.CheckCollisionCircles(e.position, 15, p.position, p.radius)
}

func (p *Projectile) Update() {
	p.position = rl.Vector2Add(p.position, p.velocity)
}

func (p *Projectile) Draw() {
	rl.DrawCircleV(p.position, p.radius, rl.RayWhite)

}

// returns direction by subtracting end-start
func GetDirection(start, end rl.Vector2) rl.Vector2 {
	return rl.NewVector2(end.X-start.X, end.Y-start.Y) // Vector A->B = B-A
}

// returns normalized direction - useful for distance
func GetDirectionNorm(start, end rl.Vector2) rl.Vector2 {
	dir := rl.NewVector2(end.X-start.X, end.Y-start.Y)
	return rl.Vector2Normalize(dir)
}

func GetMousePosNormal() rl.Vector2 {
	mousePos := rl.GetMousePosition()
	dir := GetDirection(player.position, mousePos)
	return rl.Vector2Normalize(dir)
}

func GetMousePosOffset(offset float32) rl.Vector2 {
	direction := GetMousePosNormal()
	result := rl.NewVector2(player.position.X+direction.X*offset, player.position.Y+direction.Y*offset)
	return result
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

// Get the angle in degrees for rotation
func GetAngle(v rl.Vector2) float32 {
	return float32(math.Atan2(float64(v.Y), float64(v.X)) * (180.0 / math.Pi))
}
