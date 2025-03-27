package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// type CollisionInfo struct {
// 	EntityA, EntityB Entity
// 	MTV              Vector2 // The smallest separating axis
// 	Normal           Vector2 // The collision normal
// 	Penetration      float32 // How much they overlap
// }

// Polygon represents a convex shape with a set of points.
type Polygon struct {
	Vertices []rl.Vector2
}

// NewPolygon: Creates and returns a new pointer to a Polygon
// composed of the passed in points.  Points are
// considered to be in order such that the last point
// forms an edge with the first point.
// credit - https://github.com/kellydunn/golang-geo/blob/master/polygon.go
func NewPolygon(points []rl.Vector2) *Polygon {
	return &Polygon{Vertices: points}
}

// Add: Appends the passed in contour to the current Polygon.
func (p *Polygon) Add(vertex rl.Vector2) {
	p.Vertices = append(p.Vertices, vertex)
}

func (p *Polygon) Draw() {
	for i, v := range p.Vertices {

		if i == len(p.Vertices)-1 {
			rl.DrawCircleV(v, 2, rl.Orange)
			rl.DrawLineEx(v, p.Vertices[0], 3, rl.Green)
			break
		}
		rl.DrawCircleV(v, 2, rl.Orange)
		rl.DrawLineEx(v, p.Vertices[i+1], 2, rl.Green)
	}
	rl.DrawCircleV(p.Centroid(), 2, rl.Red)
}

func (p *Polygon) SignedArea() float32 {
	sum := float32(0)
	for i, v := range p.Vertices {
		sum += v.X*p.Vertices[i+1].Y - p.Vertices[i+1].X*v.Y
	}
	return sum / 2
}

// Geometrischer Schwerpunkt
func (p *Polygon) Centroid() rl.Vector2 {
	k := len(p.Vertices)
	var sumX, sumY float32
	sumX = float32(0)
	sumY = float32(0)

	for i := 0; i <= k; i++ {
		sumX += p.Vertices[i%k].X
		sumY += p.Vertices[i%k].Y
	}
	result := rl.NewVector2(sumX/float32(k+1), sumY/float32(k+1))
	return result
}

func HandleSatCollision(p1, p2 Polygon) (float32, rl.Vector2) {
	axes := getAxes(p1)
	seperation := math.SmallestNonzeroFloat32
	minSepAxis := math.MaxFloat32
	idx := -1

	for i, va := range p1.Vertices {
		minSep := math.MaxFloat32
		for _, vb := range p2.Vertices {
			dot := rl.Vector2DotProduct(rl.Vector2Subtract(vb, va), axes[i])
			minSep = math.Min(minSep, float64(dot))
		}
		if minSep > seperation {
			seperation = minSep
			idx = i
		}
		if minSep < minSepAxis {
			minSepAxis = minSep
		}
	}
	return float32(seperation), axes[idx]
}

// SATCollision checks if two convex polygons are colliding using the Separating Axis Theorem (SAT).
func SatCollision(p1, p2 Polygon) bool {
	// Get all unique axes to test (normals of each polygon's edges)
	axes := getAxes(p1)
	axes = append(axes, getAxes(p2)...)

	// Test projections on each axis
	for _, axis := range axes {
		min1, max1 := projectPolygon(p1, axis)
		min2, max2 := projectPolygon(p2, axis)

		// If projections do not overlap, no collision / not a seperating axis
		if max1 < min2 || max2 < min1 {
			return false
		}
	}

	// If all projections overlap, collision detected
	return true
}

// getAxes returns a list of perpendicular axes from a polygon’s edges.
func getAxes(p Polygon) []rl.Vector2 {
	var axes []rl.Vector2
	for i := 0; i < len(p.Vertices); i++ {
		// Get current and next vertex
		current := p.Vertices[i]
		next := p.Vertices[(i+1)%len(p.Vertices)]

		// Edge vector
		edge := rl.Vector2Subtract(next, current)

		// Normal (perpendicular vector)
		normal := rl.Vector2{X: -edge.Y, Y: edge.X}

		// Normalize the axis
		normal = rl.Vector2Normalize(normal)
		axes = append(axes, normal)
	}
	return axes
}

// projectPolygon projects a polygon onto an axis and returns the min/max values.
func projectPolygon(p Polygon, axis rl.Vector2) (float32, float32) {
	// Project first vertex
	dot := rl.Vector2DotProduct(p.Vertices[0], axis)
	min, max := dot, dot

	// Project remaining Vertices
	for i := 1; i < len(p.Vertices); i++ {
		dot = rl.Vector2DotProduct(p.Vertices[i], axis)
		if dot < min {
			min = dot
		}
		if dot > max {
			max = dot
		}
	}
	return min, max
}

// Old ##################################
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
