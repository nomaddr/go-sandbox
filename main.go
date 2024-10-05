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

func (s *Segment) Update(prev Segment) {
	s.angle = math.Atan2(prev.y-s.y, prev.x-s.x)
	d := math.Sqrt(math.Pow(prev.x-s.x, 2) + math.Pow(prev.y-s.y, 2))
	if d > s.distance {
		delta := d - s.distance
		s.x += delta * math.Cos(s.angle)
		s.y += delta * math.Sin(s.angle)
	}
}

func (s *Segment) Draw() {
	//rl.DrawCircleLines(int32(s.x), int32(s.y), s.radius, s.color)
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

func (c *Creature) Update() {
	c.body[0].angle = math.Atan2(rl.GetMousePosition().Y-c.body[0].y, rl.GetMousePosition().X-c.body[0].x)
	c.body[0].x += c.speed * math.Cos(c.body[0].angle)
	c.body[0].y += c.speed * math.Sin(c.body[0].angle)
	for i := int32(1); i < c.len; i++ {
		c.body[i].Update(c.body[i-1])
	}
}

func (c *Creature) Draw() {
	for i := int32(0); i < c.len; i++ {
		seg := c.body[i]
		seg.Draw()
		if i == 0 {
			continue
		}
		prev := c.body[i-1]
		leftX1 := seg.x + seg.radius*math.Cos(seg.angle-0.5*math.Pi)
		leftY1 := seg.y + seg.radius*math.Sin(seg.angle-0.5*math.Pi)
		leftX2 := prev.x + prev.radius*math.Cos(prev.angle-0.5*math.Pi)
		leftY2 := prev.y + prev.radius*math.Sin(prev.angle-0.5*math.Pi)
		rl.DrawLine(int32(leftX1), int32(leftY1), int32(leftX2), int32(leftY2), rl.Purple)

		rightX1 := seg.x + seg.radius*math.Cos(seg.angle+0.5*math.Pi)
		rightY1 := seg.y + seg.radius*math.Sin(seg.angle+0.5*math.Pi)
		rightX2 := prev.x + prev.radius*math.Cos(prev.angle+0.5*math.Pi)
		rightY2 := prev.y + prev.radius*math.Sin(prev.angle+0.5*math.Pi)
		rl.DrawLine(int32(rightX1), int32(rightY1), int32(rightX2), int32(rightY2), rl.Purple)
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
		creature.Update()
		//----------------------------------------------------------------------------------
		// Draw
		//----------------------------------------------------------------------------------
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		creature.Draw()

		rl.DrawFPS(10, 10)

		rl.EndDrawing()
		//----------------------------------------------------------------------------------
	}

	// De-Initialization
	//--------------------------------------------------------------------------------------
	rl.CloseWindow() // Close window and OpenGL context
	//--------------------------------------------------------------------------------------
}
