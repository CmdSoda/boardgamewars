package hexagon

func contains(hexagonList []HexPosition, hex HexPosition) bool {
	for _, h := range hexagonList {
		if h.Equal(hex) {
			return true
		}
	}
	return false
}

func GetNextPosition(currentPosition HexPosition, path []HexPosition) HexPosition {
	for i := 0; i < len(path); i++ {
		if path[i] == currentPosition {
			return path[i+1]
		}
	}
	return currentPosition
}

// CalculatePath berechnet die direkten Pfad (LOS) zwischen zwei Hexagons.
func CalculatePath(startHex HexPosition, endHex HexPosition) []HexPosition {
	var path = make([]HexPosition, 0)
	var ignore = make([]HexPosition, 0)
	var current = startHex
	var next = HexPosition{}

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
