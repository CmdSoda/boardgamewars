package game

import (
	"github.com/CmdSoda/boardgamewars/internal/randomizer"
	"testing"
)

func TestDogfight(t *testing.T) {
	randomizer.Init()
	LoadAircrafts()
	LoadAir2AirWeapons()

	fighter1 := NewAircraftByName("F14", "Default")
	fighter2 := NewAircraftByName("F14", "Default")

	Dogfight(fighter1, fighter2)
}
