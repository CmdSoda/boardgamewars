package hexagon

import "github.com/CmdSoda/boardgamewars/internal/math"

type hexagon struct {
	column int
	row    int
}

// GetCenterCoordinates ermittelt den Mittelpunkt eines hexagons.
func GetCenterCoordinates(h hexagon) math.Vector2 {
	center := math.Vector2{}

	if h.row%2 == 0 {
		center.X = hr + float64(h.column)*float64(2)*hr
		center.Y = 1 + 1.5*float64(h.row)
	} else {
		center.X = 2*hr + 2*hr*float64(h.column)
		center.Y = 2.5 + (float64(h.row)-1)*1.5
	}

	return center
}
