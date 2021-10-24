package game

import (
	"github.com/CmdSoda/boardgamewars/internal/countrycodes"
	"github.com/CmdSoda/boardgamewars/internal/military"
	"github.com/CmdSoda/boardgamewars/internal/randomizer"
	"testing"
)

func TestDogfight(t *testing.T) {
	randomizer.Init()
	LoadAircrafts()
	LoadAir2AirWeapons()

	fighter1 := NewAircraftByName("F14", "Default", countrycodes.UK, military.NatoOfficerCodeOF2)
	fighter2 := NewAircraftByName("F14", "Default", countrycodes.UK, military.NatoOfficerCodeOF1)

	Dogfight(fighter1, fighter2)
}