package bgw

import "testing"

func TestDogfight(t *testing.T) {
	fighter1 := Aircraft{
		Type: F14,
	}
	fighter2 := Aircraft{
		Type: Mig23,
	}
	Dogfight([]Aircraft{fighter1}, []Aircraft{fighter2})
}
