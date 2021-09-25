package segment

import (
	"github.com/CmdSoda/boardgamewars/internal/math"
)

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

func NewSegment(direction Direction, center math.Vector2) Segment {
	switch direction {
	case NW:
		return Segment{
			Start: math.Vector2{X: center.X - math.Hr, Y: center.Y - 0.5},
			End:   math.Vector2{X: center.X, Y: center.Y - 1},
		}
	case NE:
		return Segment{
			Start: math.Vector2{X: center.X, Y: center.Y - 1},
			End:   math.Vector2{X: center.X + math.Hr, Y: center.Y - 0.5},
		}
	case E:
		return Segment{
			Start: math.Vector2{X: center.X + math.Hr, Y: center.Y - 0.5},
			End:   math.Vector2{X: center.X + math.Hr, Y: center.Y + 0.5},
		}
	case SE:
		return Segment{
			Start: math.Vector2{X: center.X + math.Hr, Y: center.Y + 0.5},
			End:   math.Vector2{X: center.X + math.Hr, Y: center.Y + 1},
		}
	case SW:
		return Segment{
			Start: math.Vector2{X: center.X, Y: center.Y + 1},
			End:   math.Vector2{X: center.X - math.Hr, Y: center.Y + 0.5},
		}
	case W:
		return Segment{
			Start: math.Vector2{X: center.X - math.Hr, Y: center.Y + 0.5},
			End:   math.Vector2{X: center.X - math.Hr, Y: center.Y - 0.5},
		}
	}
	return Segment{}
}
