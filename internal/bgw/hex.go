package bgw

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/vector"
	"github.com/CmdSoda/boardgamewars/internal/vector3"
	"math"
)

// HexWidth breite eines Position.
var HexWidth = math.Sqrt(3)

// HexHalfWidth halbe breite eines Position.
var HexHalfWidth = HexWidth / 2

type Position struct {
	Column int
	Row    int
}

func NewHexagon(c int, r int) Position {
	return Position{
		Column: c,
		Row:    r,
	}
}

func (h Position) String() string {
	return fmt.Sprintf("(%d, %d)", h.Column, h.Row)
}

// GetCenterCoordinates ermittelt den Mittelpunkt eines hexagons.
func (h Position) GetCenterCoordinates(useOffset bool) vector.Vector {
	center := vector.Vector{}

	var offsetValue = 0.0
	if useOffset {
		offsetValue = 0.1
	}

	if h.Row%2 == 0 {
		center.X = 2*HexHalfWidth + 2*HexHalfWidth*float64(h.Column) + offsetValue
		center.Y = 2.5 + float64(h.Row-1)*1.5 + offsetValue
	} else {
		center.X = HexHalfWidth + float64(h.Column)*2*HexHalfWidth + offsetValue
		center.Y = 1 + 1.5*float64(h.Row) + offsetValue
	}

	return center
}

func (h Position) GetSegments() [6]Segment {
	segments := [6]Segment{}
	for i := 0; i < len(segments); i++ {
		segments[i] = NewSegment(Direction(i), h.GetCenterCoordinates(false))
	}
	return segments
}

func (h Position) GetAdjacent(direction Direction) *Position {
	retHex := Position{}
	switch direction {
	case NW:
		if h.Row%2 == 0 {
			retHex = Position{
				Column: h.Column,
				Row:    h.Row + 1,
			}
		} else {
			retHex = Position{
				Column: h.Column - 1,
				Row:    h.Row + 1,
			}
		}
	case NE:
		if h.Row%2 == 0 {
			retHex = Position{
				Column: h.Column + 1,
				Row:    h.Row + 1,
			}
		} else {
			retHex = Position{
				Column: h.Column,
				Row:    h.Row + 1,
			}
		}
	case E:
		retHex = Position{
			Column: h.Column + 1,
			Row:    h.Row,
		}
	case SE:
		if h.Row%2 == 0 {
			retHex = Position{
				Column: h.Column + 1,
				Row:    h.Row - 1,
			}
		} else {
			retHex = Position{
				Column: h.Column,
				Row:    h.Row - 1,
			}
		}
	case SW:
		if h.Row%2 == 0 {
			retHex = Position{
				Column: h.Column,
				Row:    h.Row - 1,
			}
		} else {
			retHex = Position{
				Column: h.Column - 1,
				Row:    h.Row - 1,
			}
		}
	case W:
		retHex = Position{
			Column: h.Column - 1,
			Row:    h.Row,
		}
	}

	if &retHex == nil || retHex.Column <= 0 || retHex.Row <= 0 {
		return nil
	}

	return &retHex
}

func (h Position) Equal(h2 Position) bool {
	return h.Row == h2.Row && h.Column == h2.Column
}

func (h Position) CalculateIntersectionCount(startHex Position, endHex Position) int {
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

func CalculateIntersectionScalar(start1 vector.Vector, direction1 vector.Vector, start2 vector.Vector, direction2 vector.Vector) float64 {
	return (direction1.Y*start1.X - direction1.X*start1.Y - direction1.Y*start2.X + direction1.X*start2.Y) /
		(direction1.Y*direction2.X - direction1.X*direction2.Y)
}

func CalculateDistancePointLine(point vector.Vector, lineStart vector.Vector, lineEnd vector.Vector) float64 {
	p3 := vector3.NewVector(point.X, point.Y, 0)
	start3 := vector3.NewVector(lineStart.X, lineStart.Y, 0)
	end3 := vector3.NewVector(lineEnd.X, lineEnd.Y, 0)
	b := end3.Sub(start3)
	return b.Cross(p3.Sub(start3)).Norm() / b.Norm()
}

type PositionList []Position

func (poslist PositionList) Contains(p Position) bool {
	for _, p2 := range poslist {
		if p2.Equal(p) {
			return true
		}
	}
	return false
}

func (poslist PositionList) LastElement() Position {
	return poslist[len(poslist)-1]
}

func (poslist PositionList) String() string {
	ret := fmt.Sprintf("[%d]", len(poslist))
	for _, p := range poslist {
		ret = ret + fmt.Sprintf("(%d %d) ", p.Column, p.Row)
	}
	return ret
}
