package hexagon

type Segment struct {
	Start Vector2
	End   Vector2
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

func CreateSegment(direction Direction, center Vector2) Segment {
	switch direction {
	case NW:
		return Segment{
			Start: Vector2{center.X - hr, center.Y - 0.5},
			End:   Vector2{center.X, center.Y - 1},
		}
	case NE:
		return Segment{
			Start: Vector2{center.X, center.Y - 1},
			End:   Vector2{center.X + hr, center.Y - 0.5},
		}
	}
	return Segment{}
}
