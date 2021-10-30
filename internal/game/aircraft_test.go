package game

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAircraftPilots(t *testing.T) {
	assert.Nil(t, InitGame())
	ac := NewAircraft("F14", "Default", &WarPartyIdUK)
	fmt.Println(ac)
	ac = NewAircraft("F14", "Default", &WarPartyIdGermany)
	fmt.Println(ac)
	ac = NewAircraft("F14", "Default", &WarPartyIdUSA)
	fmt.Println(ac)
	ac = NewAircraft("F14", "Default", &WarPartyIdRussia)
	fmt.Println(ac)
}

func TestAssign(t *testing.T) {
	assert.Nil(t, InitGame())
	wrongid := AirbaseId(uuid.New())
	ac := NewAircraft("F14", "Default", &WarPartyIdUK)
	assert.Equal(t, false, ac.AssignToAB(&wrongid))
	nellis := NewAirbase("Nellis AB", WarPartyIdUSA, Position{6, 9})
	assert.Equal(t, true, ac.AssignToAB(&nellis.Id))
}

func TestAircraftMap(t *testing.T) {
	assert.Nil(t, InitGame())
	ab := NewAirbase("Carrier", WarPartyIdUK, Position{4, 6})
	ac := NewAircraft("F14", "Default", &WarPartyIdUK)
	ab.AddToHangar(ac)
	ac.Damage = append(ac.Damage, DamageTypeCockpit)
	id := ac.AircraftId
	fmt.Println(ab.AircraftsHangar[id])
}
