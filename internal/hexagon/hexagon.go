package hexagon

import (
	"github.com/CmdSoda/boardgamewars/internal/math"
	"github.com/CmdSoda/boardgamewars/internal/segment"
)

type hexagon struct {
	column int
	row    int
}

// GetCenterCoordinates ermittelt den Mittelpunkt eines hexagons.
func GetCenterCoordinates(h hexagon) math.Vector2 {
	center := math.Vector2{}

	if h.row%2 == 0 {
		center.X = math.Hr + float64(h.column)*float64(2)*math.Hr
		center.Y = 1 + 1.5*float64(h.row)
	} else {
		center.X = 2*math.Hr + 2*math.Hr*float64(h.column)
		center.Y = 2.5 + (float64(h.row)-1)*1.5
	}

	return center
}

func (h hexagon) GetSegments() [6]segment.Segment {
	segments := [6]segment.Segment{}
	for i := 0; i < len(segments); i++ {
		segments[i] = segment.NewSegment(segment.Direction(i), GetCenterCoordinates(h))
	}
	return segments
}

func (h hexagon) GetAdjacent(direction segment.Direction) hexagon {
	switch direction {
	case segment.NW:
		if h.row%2 == 0 {
			return hexagon{
				column: h.column - 1,
				row:    h.row - 1,
			}
		} else {
			return hexagon{
				column: h.column,
				row:    h.row - 1,
			}
		}
	case segment.NE:
		if h.row%2 == 0 {
			return hexagon{
				column: h.column,
				row:    h.row - 1,
			}
		} else {
			return hexagon{
				column: h.column + 1,
				row:    h.row - 1,
			}
		}
	case segment.E:
		if h.row%2 == 0 {
			return hexagon{
				column: h.column + 1,
				row:    h.row,
			}
		}
	case segment.SE:
		if h.row%2 == 0 {
			return hexagon{
				column: h.column,
				row:    h.row + 1,
			}
		} else {
			return hexagon{
				column: h.column + 1,
				row:    h.row + 1,
			}
		}
	case segment.SW:
		if h.row%2 == 0 {
			return hexagon{
				column: h.column - 1,
				row:    h.row + 1,
			}
		} else {
			return hexagon{
				column: h.column,
				row:    h.row + 1,
			}
		}
	case segment.W:
		if h.row%2 == 0 {
			return hexagon{
				column: h.column - 1,
				row:    h.row,
			}
		}
	default:
		return hexagon{}
	}
	return hexagon{}
}
