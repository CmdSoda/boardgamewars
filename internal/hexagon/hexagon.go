package hexagon

type hexagon struct {
	column int
	row    int
}

// GetCenterCoordinates ermittelt den Mittelpunkt eines hexagons.
func GetCenterCoordinates(h hexagon) Vector2 {
	center := Vector2{}

	if h.row%2 == 0 {
		center.x = hr + float64(h.column)*float64(2)*hr
		center.y = 1 + 1.5*float64(h.row)
	} else {
		center.x = 2*hr + 2*hr*float64(h.column)
		center.y = 2.5 + (float64(h.row)-1)*1.5
	}

	return center
}
