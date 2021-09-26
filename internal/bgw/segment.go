package bgw

import (
	"github.com/CmdSoda/boardgamewars/internal/vector"
)

type Segment struct {
	Start vector.Vector
	End   vector.Vector
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

func NewSegment(direction Direction, center vector.Vector) Segment {
	switch direction {
	case NW:
		return Segment{
			Start: vector.Vector{X: center.X - HexHalfWidth, Y: center.Y - 0.5},
			End:   vector.Vector{X: center.X, Y: center.Y - 1},
		}
	case NE:
		return Segment{
			Start: vector.Vector{X: center.X, Y: center.Y - 1},
			End:   vector.Vector{X: center.X + HexHalfWidth, Y: center.Y - 0.5},
		}
	case E:
		return Segment{
			Start: vector.Vector{X: center.X + HexHalfWidth, Y: center.Y - 0.5},
			End:   vector.Vector{X: center.X + HexHalfWidth, Y: center.Y + 0.5},
		}
	case SE:
		return Segment{
			Start: vector.Vector{X: center.X + HexHalfWidth, Y: center.Y + 0.5},
			End:   vector.Vector{X: center.X, Y: center.Y + 1},
		}
	case SW:
		return Segment{
			Start: vector.Vector{X: center.X, Y: center.Y + 1},
			End:   vector.Vector{X: center.X - HexHalfWidth, Y: center.Y + 0.5},
		}
	case W:
		return Segment{
			Start: vector.Vector{X: center.X - HexHalfWidth, Y: center.Y + 0.5},
			End:   vector.Vector{X: center.X - HexHalfWidth, Y: center.Y - 0.5},
		}
	}
	return Segment{}
}
