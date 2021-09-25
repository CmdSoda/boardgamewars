package path

import (
	"github.com/CmdSoda/boardgamewars/internal/hexagon"
	"github.com/CmdSoda/boardgamewars/internal/segment"
	"reflect"
)

func contains(hexagonList []hexagon.Hexagon, hex hexagon.Hexagon) bool {
	for _, h := range hexagonList {
		if reflect.DeepEqual(h, hex) {
			return true
		}
	}
	return false
}

func CalculatePath(startHex hexagon.Hexagon, endHex hexagon.Hexagon) []hexagon.Hexagon {
	var path = make([]hexagon.Hexagon, 0)
	var ignore = make([]hexagon.Hexagon, 0)
	var current = startHex
	var next = hexagon.Hexagon{}

	path = append(path, current)

	if startHex.Column == 0 || startHex.Row == 0 || endHex.Column == 0 || endHex.Row == 0 {
		return path
	}

	if startHex.Equal(endHex) {
		return path
	}

	for true {
		var minDist float64 = 0
		for nb := segment.NW; nb <= segment.W; nb++ {
			neighbor := current.GetAdjacent(nb)

			if neighbor != nil {
				if neighbor.Equal(endHex) {
					path = append(path, *neighbor)
					return path
				}
				dist := hexagon.CalculateDistancePointLine(neighbor.GetCenterCoordinates(true), startHex.GetCenterCoordinates(true), endHex.GetCenterCoordinates(true))
				intersectionCount := neighbor.CalculateIntersectionCount(startHex, endHex)

				if !contains(path, *neighbor) && !contains(ignore, *neighbor) && intersectionCount > 0 && (dist < minDist || minDist == 0) && neighbor.Column != 0 && neighbor.Row != 0 {
					minDist = dist
					next = *neighbor
				}
			}
		}

		path = append(path, next)

		for nb := segment.NW; nb <= segment.W; nb++ {
			neighbor := current.GetAdjacent(nb)

			if neighbor.Equal(next) == false {
				if contains(ignore, *neighbor) == false {
					ignore = append(ignore, *neighbor)
				}
			}
		}
		current = next
	}

	return path
}
