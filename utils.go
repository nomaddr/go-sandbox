package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Get the direction vector from start to end
func GetDirection(start, end rl.Vector2) rl.Vector2 {
	return rl.Vector2Subtract(end, start)
}

// Get a normalized direction vector
func GetDirectionNorm(start, end rl.Vector2) rl.Vector2 {
	dir := rl.Vector2Subtract(end, start)
	return rl.Vector2Normalize(dir)
}

// Get a vector pointing in the direction of the mouse from the player
func GetMousePosNormal() rl.Vector2 {
	return GetDirectionNorm(player.position, rl.GetMousePosition())
}

// Get an offset position in the direction of the mouse
func GetMousePosOffset(offset float32) rl.Vector2 {
	dir := GetMousePosNormal()
	return rl.Vector2Add(player.position, rl.Vector2Scale(dir, offset))
}

// Get the angle (in degrees) of a vector
func GetAngle(v rl.Vector2) float32 {
	return float32(math.Atan2(float64(v.Y), float64(v.X)) * (180.0 / math.Pi))
}
