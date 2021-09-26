package bgw

func contains(hexagonList []Position, hex Position) bool {
	for _, h := range hexagonList {
		if h.Equal(hex) {
			return true
		}
	}
	return false
}

// CalculatePath berechnet die direkten Pfad (LOS) zwischen zwei Hexagons.
func CalculatePath(startHex Position, endHex Position) []Position {
	var path = make([]Position, 0)
	var ignore = make([]Position, 0)
	var current = startHex
	var next = Position{}

	path = append(path, current)

	if startHex.Column == 0 || startHex.Row == 0 || endHex.Column == 0 || endHex.Row == 0 {
		return path
	}

	if startHex.Equal(endHex) {
		return path
	}

	for true {
		var minDist float64 = 0
		for nb := NW; nb <= W; nb++ {
			neighbor := current.GetAdjacent(nb)

			if neighbor != nil {
				if neighbor.Equal(endHex) {
					path = append(path, *neighbor)
					return path
				}
				dist := CalculateDistancePointLine(neighbor.GetCenterCoordinates(true),
					startHex.GetCenterCoordinates(true), endHex.GetCenterCoordinates(true))
				intersectionCount := neighbor.CalculateIntersectionCount(startHex, endHex)

				if !contains(path, *neighbor) && !contains(ignore, *neighbor) && intersectionCount > 0 &&
					(dist < minDist || minDist == 0) && neighbor.Column != 0 && neighbor.Row != 0 {
					minDist = dist
					next = *neighbor
				}
			}
		}

		path = append(path, next)

		for nb := NW; nb <= W; nb++ {
			neighbor := current.GetAdjacent(nb)

			if neighbor != nil {
				if neighbor.Equal(next) == false {
					if contains(ignore, *neighbor) == false {
						ignore = append(ignore, *neighbor)
					}
				}
			}
		}
		current = next
	}

	return path
}
