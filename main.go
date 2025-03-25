package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	Init()

	for !rl.WindowShouldClose() {
		Update()
		Draw()
	}

	rl.CloseWindow()
}
