package hexagon

import "github.com/CmdSoda/boardgamewars/internal/math"

type Segment struct {
	Start math.Vector2
	End   math.Vector2
}

type Direction int

const (
	NW Direction = iota
	NE
	E
	SE
	SW
	W
)

func CreateSegment(direction Direction, center math.Vector2) Segment {
	switch direction {
	case NW:
		return Segment{
			Start: math.Vector2{center.X - hr, center.Y - 0.5},
			End:   math.Vector2{center.X, center.Y - 1},
		}
	case NE:
		return Segment{
			Start: math.Vector2{center.X, center.Y - 1},
			End:   math.Vector2{center.X + hr, center.Y - 0.5},
		}
	}
	return Segment{}
}
