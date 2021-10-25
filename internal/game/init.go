package game

import "github.com/CmdSoda/boardgamewars/internal/randomizer"

func InitGame() {
	NewAirbaseList()
	NewPilotRoster()
	randomizer.Init()
	LoadAircrafts()
	LoadAir2AirWeapons()
}
