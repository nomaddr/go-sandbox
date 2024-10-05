package main

import (
	math "github.com/chewxy/math32"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	screenWidth  int32 = 800
	screenHeight int32 = 450
)

type Segment struct {
	x, y                    float32
	angle, distance, radius float32
	color                   rl.Color
}

type Creature struct {
	body  []Segment
	len   int32
	speed float32
}

func (s *Segment) update(prev Segment) {
	s.angle = math.Atan2(prev.y-s.y, prev.x-s.x)
	d := math.Sqrt(math.Pow(prev.x-s.x, 2) + math.Pow(prev.y-s.y, 2))
	if d > s.distance {
		delta := d - s.distance
		s.x += delta * math.Cos(s.angle)
		s.y += delta * math.Sin(s.angle)
	}
}

func (s *Segment) draw() {
	rl.DrawCircleLines(int32(s.x), int32(s.y), s.radius, s.color)
	endPos := rl.NewVector2(s.x+s.distance*math.Cos(s.angle), s.y+s.distance*math.Sin(s.angle))
	rl.DrawLineV(rl.NewVector2(s.x, s.y), endPos, s.color) //rl.NewVector2(s.x+s.radius, s.y+s.radius), s.color)
}

// Creature
func NewCreature(length int32, speed, radius float32) *Creature {
	var c Creature
	c.len = length
	c.body = make([]Segment, length)
	c.speed = speed
	r1 := radius
	for i := 0; i < int(c.len); i++ {
		r := r1 - float32(i)*(r1/(float32(length)-1))
		c.body[i] = Segment{float32(screenWidth)/2 - float32(i)*r1, float32(screenHeight) / 2, 0, r, r, rl.RayWhite}
	}
	return &c
}

func (c *Creature) update() {
	c.body[0].angle = math.Atan2(rl.GetMousePosition().Y-c.body[0].y, rl.GetMousePosition().X-c.body[0].x)
	c.body[0].x += c.speed * math.Cos(c.body[0].angle)
	c.body[0].y += c.speed * math.Sin(c.body[0].angle)
	for i := int32(1); i < c.len; i++ {
		c.body[i].update(c.body[i-1])
	}
}

func (c *Creature) draw() {
	for _, seg := range c.body {
		seg.draw()
	}
}

func main() {
	// Initialization
	//--------------------------------------------------------------------------------------

	rl.InitWindow(screenWidth, screenHeight, "Food and Mouse Game")

	rl.SetTargetFPS(60) // Set our game to run at 60 frames-per-second
	//--------------------------------------------------------------------------------------

	creature := NewCreature(20, 2, 20)

	// Main game loop
	for !rl.WindowShouldClose() { // Detect window close button or ESC key
		// Update
		//----------------------------------------------------------------------------------
		creature.update()
		//----------------------------------------------------------------------------------
		// Draw
		//----------------------------------------------------------------------------------
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		creature.draw()

		rl.DrawFPS(10, 10)

		rl.EndDrawing()
		//----------------------------------------------------------------------------------
	}

	// De-Initialization
	//--------------------------------------------------------------------------------------
	rl.CloseWindow() // Close window and OpenGL context
	//--------------------------------------------------------------------------------------
}
