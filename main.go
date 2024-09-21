package main

import (
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const ()

// Game type
var (
	ScreenWidth   int32
	ScreenHeight  int32
	Cols          int32
	Rows          int32
	FramesCounter int32
	Playing       bool
)

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	Init()

	rl.InitWindow(ScreenWidth, ScreenHeight, "sandbox")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		if Playing {
			Update()
		}

		Input()

		Draw()
	}

	rl.CloseWindow()
}

// Init - Initialize game
func Init() {
	ScreenWidth = 800
	ScreenHeight = 450
	FramesCounter = 0

}

// Input - Game input
func Input() {
	// control
	if rl.IsKeyPressed(rl.KeyD) {
	}
	if rl.IsKeyPressed(rl.KeyA) {
	}
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
	}
	if rl.IsKeyPressed(rl.KeySpace) {
	}

	FramesCounter++
}

// Update - Update game
func Update() {

}

// Draw - Draw game
func Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	rl.EndDrawing()
}
