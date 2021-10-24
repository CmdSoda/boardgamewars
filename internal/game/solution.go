package game

import "github.com/CmdSoda/boardgamewars/internal/hexagon"

type Solution struct {
	Path hexagon.PositionList
}

type SolutionList []Solution
