package bgw

import "testing"

func TestDogfight(t *testing.T) {
	RandInit()
	LoadAircrafts()
	LoadAir2AirWeapons()

	fighter1 := NewAircraftByName("F14", "Default")
	fighter2 := NewAircraftByName("F14", "Default")

	Dogfight(fighter1, fighter2)
}
