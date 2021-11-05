package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/nato"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAircraftPilots(t *testing.T) {
	assert.Nil(t, InitGame())
	ac := NewAircraft("F14", "Default", WarPartyIdUK)
	fmt.Println(ac)
	ac = NewAircraft("F14", "Default", WarPartyIdGermany)
	fmt.Println(ac)
	ac = NewAircraft("F14", "Default", WarPartyIdUSA)
	fmt.Println(ac)
	ac = NewAircraft("F14", "Default", WarPartyIdRussia)
	fmt.Println(ac)
}

func TestAssign(t *testing.T) {
	assert.Nil(t, InitGame())
	wrongid := AirbaseId(uuid.New())
	ac := NewAircraft("F14", "Default", WarPartyIdUK)
	assert.Equal(t, false, ac.AssignToAB(wrongid))
	nellis := NewAirbase("Nellis AB", WarPartyIdUSA, FloatPosition{6, 9})
	assert.Equal(t, true, ac.AssignToAB(nellis.AirbaseId))
}

func TestAircraftMap(t *testing.T) {
	assert.Nil(t, InitGame())
	ab := NewAirbase("Carrier", WarPartyIdUK, FloatPosition{4, 6})
	ac := NewAircraft("F14", "Default", WarPartyIdUK)
	p1 := NewPilot(WarPartyIdUK, nato.OF1)
	ac.AddPilot(p1)
	p2 := NewPilot(WarPartyIdUK, nato.OF1)
	ac.AddPilot(p2)
	ab.AddToHangar(ac.AircraftId)
	ac.Damage = append(ac.Damage, DamageTypeCockpit)
	fmt.Println(ab.AircraftsHangar)
}

func TestAircraftId(t *testing.T) {
	assert.Nil(t, InitGame())
	ac := NewAircraft("F14", "Default", WarPartyIdUK)
	_, exist := Globals.AllAircrafts[ac.AircraftId]
	assert.True(t, exist)
}

func TestAircraft_AssignToAB(t *testing.T) {
	assert.Nil(t, InitGame())
	ab := NewAirbase("Carrier", WarPartyIdUK, FloatPosition{4, 6})
	ac := NewAircraft("F14", "Default", WarPartyIdUK)
	ac.AssignToAB(ab.AirbaseId)
	_, exist := Globals.AllAirbases[ab.AirbaseId]
	assert.True(t, exist)
	assert.Equal(t, ab.AirbaseId, ac.StationedAt)
}
