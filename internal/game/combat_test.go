package game

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDogfight(t *testing.T) {
	assert.Nil(t, InitGame())

	fighter1 := NewAircraft("F14", "Default", WarPartyIdUK)
	fighter2 := NewAircraft("MiG-29", "Default", WarPartyIdUK)

	assert.NotNil(t, fighter1)
	assert.NotNil(t, fighter2)

	ldp1 := DogfightPositionTossup
	ldp2 := DogfightPositionTossup
	dr1, dr2 := ExecuteDogfight(fighter1.AircraftId, ldp1, fighter2.AircraftId, ldp2)
	assert.NotNil(t, dr1)
	assert.NotNil(t, dr2)
	fmt.Println(dr1)
	fmt.Println(dr2)
}

func TestMoreRounds(t *testing.T) {
	assert.Nil(t, InitGame())

	fighter1 := NewAircraft("F14", "Default", WarPartyIdUK)
	fighter2 := NewAircraft("MiG-29", "Default", WarPartyIdRussia)

	assert.NotNil(t, fighter1)
	assert.NotNil(t, fighter2)

	drl1, drl2 := Sim10Rounds(fighter1.AircraftId, fighter2.AircraftId)
	assert.NotNil(t, drl1)
	assert.NotNil(t, drl2)

	fmt.Println(drl1)
	fmt.Println(drl2)
}
