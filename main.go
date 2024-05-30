package main

import (
	"math/rand"
	"time"

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

type Enemy struct {
	Head        rl.Vector2
	Segments    []Segment
	Speed       float32
	SegmentDist float32
}

func NewEnemy(numSegments int, speed float32, segmentDist float32) *Enemy {
	segments := make([]Segment, numSegments)
	for i := 0; i < numSegments; i++ {
		segments[i] = Segment{Position: rl.Vector2{X: -segmentDist * float32(i), Y: 0}}
	}
	return &Enemy{
		Head:        rl.Vector2{X: 400, Y: 225},
		Segments:    segments,
		Speed:       speed,
		SegmentDist: segmentDist,
	}
}

func (e *Enemy) Update(target rl.Vector2) {
	// Move the head towards the target
	if rl.Vector2Distance(e.Head, target) > e.Speed*rl.GetFrameTime() {
		e.Head = rl.Vector2MoveTowards(e.Head, target, e.Speed*rl.GetFrameTime())
	}

	// Update each segment to follow the previous one
	for i := range e.Segments {
		if i == 0 {
			// The first segment follows the head
			if rl.Vector2Distance(e.Segments[i].Position, e.Head) > e.SegmentDist {
				e.Segments[i].Position = rl.Vector2MoveTowards(e.Segments[i].Position, e.Head, e.Speed*rl.GetFrameTime())
			}
		} else {
			// Each segment follows the previous segment
			if rl.Vector2Distance(e.Segments[i].Position, e.Segments[i-1].Position) > e.SegmentDist {
				e.Segments[i].Position = rl.Vector2MoveTowards(e.Segments[i].Position, e.Segments[i-1].Position, e.Speed*rl.GetFrameTime())
			}
		}
	}
}

func (e *Enemy) Draw() {
	rl.DrawCircleV(e.Head, 20, rl.Maroon)
	for i := 0; i < len(e.Segments); i++ {
		rl.DrawCircleV(e.Segments[i].Position, 14, rl.Maroon)
		switch i {
		case 0:
			rl.DrawLineEx(e.Segments[i].Position, e.Head, 10, rl.Maroon)
		default:
			rl.DrawLineEx(e.Segments[i].Position, e.Segments[i-1].Position, 10, rl.Maroon)
		}
	}
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
)

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

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

// Init Player
func (p *Player) Init(length int) {
	p.pos = rl.NewVector2(float32(screenWidth)/2, float32(screenHeight)/2)
	p.isMoving = false
	p.speed = 4
	p.length = length

	p.segments = make([]rl.Vector2, p.length)
}

// Init - Initialize game
func Init() {
	player = Player{}
	player.Init(5)

	enemy = NewEnemy(5, 80, 50)

	mousePos = rl.Vector2{X: float32(screenWidth) / 2, Y: float32(screenHeight) / 2}

}

// Input - Game input
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
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		mousePos = rl.GetMousePosition()
	}
	if rl.IsKeyPressed(rl.KeySpace) {
	}

}

// Update - Update game
func Update() {
	enemy.Update(mousePos)
}

// Draw - Draw game
func Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.DarkGray)

	rl.DrawCircle(int32(player.pos.X), int32(player.pos.Y), 30, rl.Maroon)
	enemy.Draw()

	rl.EndDrawing()
}
