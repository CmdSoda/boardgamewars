package game

import "github.com/CmdSoda/boardgamewars/internal/randomizer"

func InitGame() error {
	var err error
	NewAirbaseList()
	NewPilotRoster()
	randomizer.Init()
	if err = LoadAircrafts(); err != nil {
		return err
	}
	if err = LoadAir2AirWeapons(); err != nil {
		return err
	}
	return nil
}
