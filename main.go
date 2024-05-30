package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 800
	screenHeight = 640
)

type Segment struct {
	Position rl.Vector2
	Speed    float32
}

type Player struct {
	pos      rl.Vector2
	dir      rl.Vector2
	segments []rl.Vector2
	isMoving bool
	speed    float32
	length   int
}

var (
	player   Player
	mousePos rl.Vector2
	enemy    *Enemy
	cam      rl.Camera2D
)

func main() {
	//	rand.New(rand.NewSource(time.Now().UnixNano()))

	//	game := Game{}
	//	game.Init(false)
	Init()

	rl.InitWindow(screenWidth, screenHeight, "sandbox")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		Input()
		Update()
		Draw()
	}

	rl.CloseWindow()
}

// ##############
// Player Methods
// ##############

func (p *Player) init(length int) {
	p.pos = rl.NewVector2(float32(screenWidth)/2, float32(screenHeight)/2)
	p.isMoving = false
	p.speed = 4
	p.length = length

}

func (p *Player) shoot() {

}

// ######################
// Init - Initialize game
// ######################

func Init() {
	player = Player{}
	player.init(5)

	enemy = NewEnemy(5, 200, 50)

	mousePos = rl.Vector2{X: float32(screenWidth) / 2, Y: float32(screenHeight) / 2}

	cam = rl.NewCamera2D(rl.NewVector2(screenWidth/2, screenHeight/2), player.pos, 0, 1)

}

// ####################
//
//	Input - Game input
//
// ####################
func Input() {
	// control
	if rl.IsKeyDown(rl.KeyW) {
		player.pos.Y -= player.speed

	}
	if rl.IsKeyDown(rl.KeyS) {
		player.pos.Y += player.speed

	}
	if rl.IsKeyDown(rl.KeyD) {
		player.pos.X += player.speed

	}
	if rl.IsKeyDown(rl.KeyA) {
		player.pos.X -= player.speed

	}
	if rl.IsMouseButtonDown(rl.MouseLeftButton) {
		mousePos = rl.GetMousePosition()
	}
	if rl.IsKeyPressed(rl.KeySpace) {

	}

}

// ######################
//
//	Update - Update game
//
// ######################
func Update() {
	enemy.update(mousePos)
}

// ##################
//
//	Draw - Draw game
//
// ##################
func Draw() {
	rl.BeginDrawing()
	rl.BeginMode2D(cam)
	rl.ClearBackground(rl.Gray)

	rect := rl.NewRectangle(player.pos.X, player.pos.Y, 10, 20)
	rl.DrawRectangleRounded(rect, 0.5, 1, rl.Purple) // rl.DrawCircle(int32(player.pos.X), int32(player.pos.Y), 10, rl.DarkPurple)
	enemy.draw()
	//drawFunc()

	rl.EndMode2D()
	rl.EndDrawing()
}
