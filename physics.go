package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	maxVertices int32 = 100
)

type Polygon struct {
	vertices []rl.Vector2
}

// NewPolygon: Creates and returns a new pointer to a Polygon
// composed of the passed in points.  Points are
// considered to be in order such that the last point
// forms an edge with the first point.
// credit - https://github.com/kellydunn/golang-geo/blob/master/polygon.go
func NewPolygon(points []rl.Vector2) *Polygon {
	return &Polygon{vertices: points}
}

// Add: Appends the passed in contour to the current Polygon.
func (p *Polygon) Add(vertex rl.Vector2) {
	p.vertices = append(p.vertices, vertex)
}

func (p *Polygon) Draw() {
	for i, v := range p.vertices {

		if i == len(p.vertices)-1 {
			rl.DrawCircleV(v, 5, rl.Orange)
			rl.DrawLineEx(v, p.vertices[0], 3, rl.Green)
			break
		}
		rl.DrawCircleV(v, 5, rl.Orange)
		rl.DrawLineEx(v, p.vertices[i+1], 3, rl.Green)
	}
}

func FindMinSeperation(a, b Polygon) float32 {
	seperation := float32(math.MinInt32)

	// for i, va := range a.vertices {
	// 	normal :=
	// }

	return seperation
}

// very simple not useful collision check
func CheckCollisions() {
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

	// Check enemy-enemy collision
	for i := 1; i < len(enemies); i++ {
		if rl.CheckCollisionCircles(enemies[i].position, 15, enemies[i-1].position, 15) {

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
