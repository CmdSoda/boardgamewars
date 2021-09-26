package hexagon

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/math"
	"github.com/CmdSoda/boardgamewars/internal/segment"
)

type Hexagon struct {
	Column int
	Row    int
}

func NewHexagon(c int, r int) Hexagon {
	return Hexagon{
		Column: c,
		Row:    r,
	}
}

func (h Hexagon) String() string {
	return fmt.Sprintf("(%d, %d)", h.Column, h.Row)
}

// GetCenterCoordinates ermittelt den Mittelpunkt eines hexagons.
func (h Hexagon) GetCenterCoordinates(useOffset bool) math.Vector2 {
	center := math.Vector2{}

	var offsetValue = 0.0
	if useOffset {
		offsetValue = 0.1
	}

	if h.Row%2 == 0 {
		center.X = 2*math.Hr + 2*math.Hr*float64(h.Column) + offsetValue
		center.Y = 2.5 + float64(h.Row-1)*1.5 + offsetValue
	} else {
		center.X = math.Hr + float64(h.Column)*2*math.Hr + offsetValue
		center.Y = 1 + 1.5*float64(h.Row) + offsetValue
	}

	return center
}

func (h Hexagon) GetSegments() [6]segment.Segment {
	segments := [6]segment.Segment{}
	for i := 0; i < len(segments); i++ {
		segments[i] = segment.NewSegment(segment.Direction(i), h.GetCenterCoordinates(false))
	}
	return segments
}

func (h Hexagon) GetAdjacent(direction segment.Direction) *Hexagon {
	retHex := Hexagon{}
	switch direction {
	case segment.NW:
		if h.Row%2 == 0 {
			retHex = Hexagon{
				Column: h.Column,
				Row:    h.Row + 1,
			}
		} else {
			retHex = Hexagon{
				Column: h.Column - 1,
				Row:    h.Row + 1,
			}
		}
	case segment.NE:
		if h.Row%2 == 0 {
			retHex = Hexagon{
				Column: h.Column + 1,
				Row:    h.Row + 1,
			}
		} else {
			retHex = Hexagon{
				Column: h.Column,
				Row:    h.Row + 1,
			}
		}
	case segment.E:
		retHex = Hexagon{
			Column: h.Column + 1,
			Row:    h.Row,
		}
	case segment.SE:
		if h.Row%2 == 0 {
			retHex = Hexagon{
				Column: h.Column + 1,
				Row:    h.Row - 1,
			}
		} else {
			retHex = Hexagon{
				Column: h.Column,
				Row:    h.Row - 1,
			}
		}
	case segment.SW:
		if h.Row%2 == 0 {
			retHex = Hexagon{
				Column: h.Column,
				Row:    h.Row - 1,
			}
		} else {
			retHex = Hexagon{
				Column: h.Column - 1,
				Row:    h.Row - 1,
			}
		}
	case segment.W:
		retHex = Hexagon{
			Column: h.Column - 1,
			Row:    h.Row,
		}
	}

	if &retHex == nil || retHex.Column <= 0 || retHex.Row <= 0 {
		return nil
	}

	return &retHex
}

func (h Hexagon) Equal(h2 Hexagon) bool {
	return h.Row == h2.Row && h.Column == h2.Column
}

func (h Hexagon) CalculateIntersectionCount(startHex Hexagon, endHex Hexagon) int {
	segments := h.GetSegments()
	count := 0
	for _, s := range segments {
		k1 := CalculateIntersectionScalar(startHex.GetCenterCoordinates(true),
			endHex.GetCenterCoordinates(true).Minus(startHex.GetCenterCoordinates(true)),
			s.Start, s.End.Minus(s.Start))
		k2 := CalculateIntersectionScalar(s.Start, s.End.Minus(s.Start), startHex.GetCenterCoordinates(true),
			endHex.GetCenterCoordinates(true).Minus(startHex.GetCenterCoordinates(true)))
		if k1 >= 0 && k1 <= 1 && k2 >= 0 && k2 <= 1 {
			count = count + 1
		}
	}
	return count
}

func CalculateIntersectionScalar(start1 math.Vector2, direction1 math.Vector2, start2 math.Vector2, direction2 math.Vector2) float64 {
	return (direction1.Y*start1.X - direction1.X*start1.Y - direction1.Y*start2.X + direction1.X*start2.Y) /
		(direction1.Y*direction2.X - direction1.X*direction2.Y)
}

func CalculateDistancePointLine(point math.Vector2, lineStart math.Vector2, lineEnd math.Vector2) float64 {
	p3 := math.NewVector3(point.X, point.Y, 0)
	start3 := math.NewVector3(lineStart.X, lineStart.Y, 0)
	end3 := math.NewVector3(lineEnd.X, lineEnd.Y, 0)
	b := end3.Sub(start3)
	return b.Cross(p3.Sub(start3)).Norm() / b.Norm()
}
