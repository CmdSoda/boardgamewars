package main

import (
	"github.com/CmdSoda/boardgamewars/internal/game"
	"github.com/sirupsen/logrus"
)

func main() {
	err := game.InitGameWithLogLevel(0, logrus.WarnLevel)
	if err != nil {
		panic("Could not init game: " + err.Error())
		return
	}
	game.FreeForAll()
	game.Globals.Statistics.AircraftVsAircraft.Dump()
}
