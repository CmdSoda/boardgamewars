package game

import "github.com/CmdSoda/boardgamewars/internal/randomizer"

func InitGame() {
	NewPilotRoster()
	randomizer.Init()
	LoadAircrafts()
	LoadAir2AirWeapons()
}
