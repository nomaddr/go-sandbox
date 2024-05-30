package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Enemy struct {
	Head        rl.Vector2
	Segments    []Segment
	Speed       float32
	SegmentDist float32
}

// #############
// Enemy Methods
// #############
func NewEnemy(numSegments int, speed float32, segmentDist float32) *Enemy {
	segments := make([]Segment, numSegments)
	for i := 0; i < numSegments; i++ {
		segments[i] = Segment{Position: rl.Vector2{X: -segmentDist * float32(i), Y: 0}}
		segments[i].Speed = speed
	}
	return &Enemy{
		Head:        rl.Vector2{X: 400, Y: 225},
		Segments:    segments,
		Speed:       speed,
		SegmentDist: segmentDist,
	}
}

func (e *Enemy) update(target rl.Vector2) {
	// Move the head towards the target
	if rl.Vector2Distance(e.Head, target) > e.Speed*rl.GetFrameTime() {
		e.Head = rl.Vector2MoveTowards(e.Head, target, e.Speed*rl.GetFrameTime())
	}

	// Update each segment to follow the previous one
	for i := range e.Segments {
		if i == 0 {
			// The first segment follows the head
			if rl.Vector2Distance(e.Segments[i].Position, e.Head) > e.SegmentDist {
				e.Segments[i].Position = rl.Vector2MoveTowards(e.Segments[i].Position, e.Head, e.Segments[i].Speed*rl.GetFrameTime())
			}
		} else {
			// Each segment follows the previous segment
			if rl.Vector2Distance(e.Segments[i].Position, e.Segments[i-1].Position) > e.SegmentDist {
				e.Segments[i].Position = rl.Vector2MoveTowards(e.Segments[i].Position, e.Segments[i-1].Position, e.Segments[i].Speed*rl.GetFrameTime())
			}
		}
	}
}

func (e *Enemy) draw() {
	rl.DrawCircleV(e.Head, 20, rl.Maroon)
	for i := 0; i < len(e.Segments); i++ {
		rl.DrawCircleV(e.Segments[i].Position, 14, rl.Maroon)
		switch i {
		case 0:
			rl.DrawLineBezier(e.Segments[i].Position, e.Head, 10, rl.DarkBlue)
		case 5:
			break
		default:
			rl.DrawLineBezier(e.Segments[i].Position, e.Segments[i-1].Position, 10, rl.DarkBlue)
		}
	}
}
