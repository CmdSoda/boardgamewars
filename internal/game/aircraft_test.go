package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/hexagon"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAircraftPilots(t *testing.T) {
	assert.Nil(t, InitGame(0))
	ac := NewAircraft("F14", "Default", "uk")
	fmt.Println(ac)
	ac = NewAircraft("F14", "Default", "germany")
	fmt.Println(ac)
	ac = NewAircraft("F14", "Default", "usa")
	fmt.Println(ac)
	ac = NewAircraft("F14", "Default", "russia")
	fmt.Println(ac)
}

func TestAssign(t *testing.T) {
	assert.Nil(t, InitGame(0))
	wrongid := AirbaseId(uuid.New())
	ac := NewAircraft("F14", "Default", "uk")
	assert.Equal(t, false, ac.AssignToAB(wrongid))
	nellis := NewAirbase("Nellis AB", "usa", hexagon.HexPosition{Column: 6, Row: 9})
	assert.Equal(t, true, ac.AssignToAB(nellis.AirbaseId))
}

func TestAircraftMap(t *testing.T) {
	assert.Nil(t, InitGame(0))
	ab := NewAirbase("Carrier", "uk", hexagon.HexPosition{Column: 4, Row: 6})
	ac := NewAircraft("F14", "Default", "uk")
	p1 := NewPilot("uk", OF1)
	ac.AddPilot(p1.PilotId)
	p2 := NewPilot("uk", OF1)
	ac.AddPilot(p2.PilotId)
	ab.AddToParkingArea(ac.AircraftId)
	ac.Damage = append(ac.Damage, DamageTypeCockpit)
	fmt.Println(ab.ParkingArea)
}

func TestAircraftId(t *testing.T) {
	assert.Nil(t, InitGame(0))
	ac := NewAircraft("F14", "Default", "uk")
	_, exist := Globals.AllAircrafts[ac.AircraftId]
	assert.True(t, exist)
}

func TestAircraft_AssignToAB(t *testing.T) {
	assert.Nil(t, InitGame(0))
	ab := NewAirbase("Carrier", "uk", hexagon.HexPosition{Column: 4, Row: 6})
	ac := NewAircraft("F14", "Default", "uk")
	ac.AssignToAB(ab.AirbaseId)
	_, exist := Globals.AllAirbases[ab.AirbaseId]
	assert.True(t, exist)
	assert.Equal(t, ab.AirbaseId, ac.StationedAt)
}

func TestAircraft_DoDamageWith(t *testing.T) {
	assert.Nil(t, InitGame(0))
	ac1 := NewAircraft("F14", "Default", "usa")
	ac1.FillSeatsWithNewPilots(OF1)
	ac2 := NewAircraft("MiG-29", "Default", "russia")
	ac2.FillSeatsWithNewPilots(OF1)
	ws, _ := ac1.GetBestDogfightingWeapon()
	ac2.DoDamageWith(ws)
}

func TestAircraft_GetHexPosition(t *testing.T) {
	assert.Nil(t, InitGameWithLogLevel(0, logrus.WarnLevel))
	ac1 := NewAircraft("F14", "Default", "usa")
	ac1.FillSeatsWithNewPilots(OF1)
	hp := ac1.GetHexPosition()
	fmt.Println(hp)
}

func TestStateChange(t *testing.T) {
	assert.Nil(t, InitGameWithLogLevel(0, logrus.WarnLevel))
	ac1 := NewAircraft("F14", "Default", "usa")
	ac1.FillSeatsWithNewPilots(OF1)

	assert.Equal(t, AcStateParking, ac1.FSM.Current())

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
	assert.Equal(t, AcStateParking, ac1.FSM.Current())
}

func TestStateChanges(t *testing.T) {
	assert.Nil(t, InitGameWithLogLevel(0, logrus.WarnLevel))
	ab := NewAirbase("Airbase1", "usa", hexagon.HexPosition{Column: 15, Row: 15})
	ab.MaxMaintenanceSlots = 4

	ac1 := NewAircraft("F14", "Default", "usa")
	ac1.FillSeatsWithNewPilots(OF1)
	assert.True(t, ac1.AssignToAB(ab.AirbaseId))

	ac1.Damage = append(ac1.Damage, []DamageType{DamageTypeFuselage, DamageTypeCockpit}...)
	ab.AddToParkingArea(ac1.AircraftId)
	err := ac1.FSM.Event(AcEventRepair)
	assert.Nil(t, err)
	assert.Equal(t, StepTime(2*20), ac1.RepairTime)

	ab.Step(10)
	ac1.Step(10)
	assert.Equal(t, AcStateInMaintenance, ac1.FSM.Current())
	ab.Step(30)
	ac1.Step(30)
	assert.Equal(t, AcStateParking, ac1.FSM.Current())
	assert.Equal(t, 0, len(ac1.Damage))
	ac1.Waypoints = []hexagon.HexPosition{{
		Column: 15,
		Row:    15,
	}}
	ab.Step(1)
	ac1.Step(1)
	assert.Equal(t, AcStateInTheAir, ac1.FSM.Current())
	ac1.Destination = hexagon.NewHexagon(15, 15)
	ac1.CurrentPosition = hexagon.NewHexagon(15, 15)
	ab.Step(1)
	ac1.Step(1)
	assert.Equal(t, AcStateParking, ac1.FSM.Current())

	fmt.Println(Globals.EventList.String())
}
