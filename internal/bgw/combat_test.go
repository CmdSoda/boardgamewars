package bgw

import "testing"

func TestDogfight(t *testing.T) {
	fighter1 := Aircraft{
		AircraftId: 0,
	}
	fighter2 := Aircraft{
		AircraftId: 0,
	}
	Dogfight(fighter1, fighter2)
}
