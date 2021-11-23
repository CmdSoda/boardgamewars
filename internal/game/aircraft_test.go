package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/nato"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAircraftPilots(t *testing.T) {
	assert.Nil(t, InitGame(0))
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
	assert.Nil(t, InitGame(0))
	wrongid := AirbaseId(uuid.New())
	ac := NewAircraft("F14", "Default", WarPartyIdUK)
	assert.Equal(t, false, ac.AssignToAB(wrongid))
	nellis := NewAirbase("Nellis AB", WarPartyIdUSA, FloatPosition{6, 9})
	assert.Equal(t, true, ac.AssignToAB(nellis.AirbaseId))
}

func TestAircraftMap(t *testing.T) {
	assert.Nil(t, InitGame(0))
	ab := NewAirbase("Carrier", WarPartyIdUK, FloatPosition{4, 6})
	ac := NewAircraft("F14", "Default", WarPartyIdUK)
	p1 := NewPilot(WarPartyIdUK, nato.OF1)
	ac.AddPilot(p1.PilotId)
	p2 := NewPilot(WarPartyIdUK, nato.OF1)
	ac.AddPilot(p2.PilotId)
	ab.AddToHangar(ac.AircraftId)
	ac.Damage = append(ac.Damage, DamageTypeCockpit)
	fmt.Println(ab.AircraftsHangar)
}

func TestAircraftId(t *testing.T) {
	assert.Nil(t, InitGame(0))
	ac := NewAircraft("F14", "Default", WarPartyIdUK)
	_, exist := Globals.AllAircrafts[ac.AircraftId]
	assert.True(t, exist)
}

func TestAircraft_AssignToAB(t *testing.T) {
	assert.Nil(t, InitGame(0))
	ab := NewAirbase("Carrier", WarPartyIdUK, FloatPosition{4, 6})
	ac := NewAircraft("F14", "Default", WarPartyIdUK)
	ac.AssignToAB(ab.AirbaseId)
	_, exist := Globals.AllAirbases[ab.AirbaseId]
	assert.True(t, exist)
	assert.Equal(t, ab.AirbaseId, ac.StationedAt)
}

func TestAircraft_DoDamageWith(t *testing.T) {
	assert.Nil(t, InitGame(0))
	ac1 := NewAircraft("F14", "Default", WarPartyIdUSA)
	ac1.FillSeatsWithNewPilots(nato.OF1)
	ac2 := NewAircraft("MiG-29", "Default", WarPartyIdRussia)
	ac2.FillSeatsWithNewPilots(nato.OF1)
	ws, _ := ac1.GetBestDogfightingWeapon()
	ac2.DoDamageWith(ws)
}

func TestAircraft_GetHexPosition(t *testing.T) {
	assert.Nil(t, InitGameWithLogLevel(0, logrus.WarnLevel))
	ac1 := NewAircraft("F14", "Default", WarPartyIdUSA)
	ac1.FillSeatsWithNewPilots(nato.OF1)
	hp := ac1.GetHexPosition()
	fmt.Println(hp)
}

func TestStateChange(t *testing.T) {
	assert.Nil(t, InitGameWithLogLevel(0, logrus.WarnLevel))
	ac1 := NewAircraft("F14", "Default", WarPartyIdUSA)
	ac1.FillSeatsWithNewPilots(nato.OF1)

	assert.Equal(t, AcStateInTheHangar, ac1.FSM.Current())

	err := ac1.FSM.Event(AcEventStart)
	assert.Nil(t, err)
	assert.Equal(t, AcStateInTheAir, ac1.FSM.Current())

	err = ac1.FSM.Event(AcEventAttack)
	assert.Nil(t, err)
	assert.Equal(t, AcStateInDogfight, ac1.FSM.Current())

	// Unerlaubt
	err = ac1.FSM.Event(AcEventLand)
	assert.NotNil(t, err)

	err = ac1.FSM.Event(AcEventDisengage)
	assert.Nil(t, err)
	assert.Equal(t, AcStateInTheAir, ac1.FSM.Current())

	err = ac1.FSM.Event(AcEventLand)
	assert.Nil(t, err)
	assert.Equal(t, AcStateInTheHangar, ac1.FSM.Current())
}

func TestRepairTime(t *testing.T) {
	assert.Nil(t, InitGameWithLogLevel(0, logrus.WarnLevel))
	ac1 := NewAircraft("F14", "Default", WarPartyIdUSA)
	ac1.FillSeatsWithNewPilots(nato.OF1)

	err := ac1.FSM.Event(AcEventRepair, StepTime(20))
	assert.Nil(t, err)
	assert.Equal(t, StepTime(20), ac1.RepairTime)
}
