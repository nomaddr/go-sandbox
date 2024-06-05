package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 800
	screenHeight = 450
)

var (
	black    = rl.NewColor(13, 0, 26, 1)
	bgX, bgY []float32
	circles  int
)

type GameObject interface {
	Update()
	Draw()
}

type Player struct {
	position rl.Vector2
	speed    float32
	health   int
	weapon   int
	timer    float32
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

type Particle struct {
	tex rl.Texture2D
	pos rl.Vector2
}

var player *Player
var enemies []*Enemy
var projectiles []*Projectile

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Top-Down Shooter")
	rl.SetTargetFPS(60)

	player = &Player{
		position: rl.NewVector2(screenWidth/2, screenHeight/2),
		speed:    200,
		health:   100,
		weapon:   1,
		timer:    0,
	}

	// Spawn initial enemies
	for i := 0; i < 5; i++ {
		enemies = append(enemies, NewEnemy())
	}
	circles = 100
	bgX, bgY = randomCoords(circles)
	for !rl.WindowShouldClose() {
		update()
		draw()
	}

	rl.CloseWindow()
}

func spawnParticle(x, y, dur, size int32, count int, pbspeed float32) {
	pbspeed = pbspeed * float32(dur)
	t := time.NewTimer(time.Duration(pbspeed))
	<-t.C
	for i := 0; i < count; i++ {
		fmt.Println(pbspeed)
		rl.DrawRectangle(x+(rand.Int31()%10), y+(rand.Int31()%10), size, size, rl.Orange)
	}
}

func NewEnemy() *Enemy {
	return &Enemy{
		position: rl.NewVector2(float32(rand.Intn(screenWidth)), float32(rand.Intn(screenHeight))),
		speed:    100,
		health:   20,
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

	if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
		if player.timer < 2 {
			player.timer += rl.GetFrameTime()
			fmt.Println(player.timer)
		}
	}
	if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
		fmt.Printf("shot: %f", player.timer)
		p.Shoot()
		player.timer = 0
	}
}

func (p *Player) Draw() {
	rl.DrawCircleV(p.position, 20, rl.Blue)
}

func (p *Player) Shoot() {
	mousePos := rl.GetMousePosition()
	direction := rl.Vector2Subtract(mousePos, p.position)
	distance := float32(math.Sqrt(float64(direction.X*direction.X + direction.Y*direction.Y)))
	var weaponSpeed float32
	switch p.weapon {
	case 0:
		weaponSpeed = 1 // no weapon
	case 1:
		weaponSpeed = 30 + player.timer*2 // charged weapon
	default:
		weaponSpeed = 30
	}
	direction = rl.Vector2Scale(direction, player.timer)
	fmt.Println(distance, " ", weaponSpeed)
	projectile := &Projectile{
		position: p.position,
		velocity: direction,
		radius:   5,
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
	rl.DrawCircleV(e.position, 15, rl.Red)
}

func (e *Enemy) IsCollidingWithProjectile(p *Projectile) bool {
	return rl.CheckCollisionCircles(e.position, 15, p.position, p.radius)
}

func (p *Projectile) Update() {
	p.position = rl.Vector2Add(p.position, p.velocity)
}

func (p *Projectile) Draw() {
	rl.DrawCircleV(p.position, p.radius, rl.Orange)
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
					spawnParticle(int32(enemies[j].position.X), int32(enemies[j].position.Y), 10, 10, 40, 1.0)
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

func drawBG(c int, x, y []float32) {
	// change radius to random for bubble effect  change time and stuff (rand.Float32()*2)+10, float32(50)*rand.Float32()
	// form circles in line: rl.DrawCircleLines(int32(x[i]/2), int32(x[i]), 2, rl.LightGray)

	for i := range c {
		rl.DrawCircleLines(int32(x[i]), int32(y[i]), 10, rl.LightGray)
		rl.DrawCircleLines(int32(x[i]/2), int32(y[i]), 2, rl.LightGray)
		rl.DrawCircle(int32(x[i]/0.2*8), int32(y[i]), float32(rand.Intn(2)), rl.Gold)
	}
}

func randomCoords(c int) ([]float32, []float32) {

	x := make([]float32, c)

	y := make([]float32, c)

	for i := range c {
		x[i] = float32(rand.Intn(screenWidth))
		y[i] = float32(rand.Intn(screenWidth))
	}

	return x, y
}

func draw() {
	rl.BeginDrawing()
	rl.ClearBackground(black)
	drawBG(circles, bgX, bgY)
	player.Draw()

	for _, enemy := range enemies {
		enemy.Draw()
	}

	for _, projectile := range projectiles {
		projectile.Draw()
	}

	// Draw player health
	rl.DrawText(fmt.Sprintf("Health: %d", player.health), 10, 10, 20, rl.RayWhite)

	rl.EndDrawing()
}
