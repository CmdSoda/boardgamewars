package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/hexagon"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewWorld(t *testing.T) {
	w := NewWorld()
	fmt.Println(w)
}

func TestStep(t *testing.T) {
	assert.Nil(t, InitGameWithLogLevel(0, logrus.WarnLevel))
	Step(3)
	assert.Equal(t, StepTime(3), Globals.World.CurrentStep)
	Step(7)
	assert.Equal(t, StepTime(10), Globals.World.CurrentStep)
}

func TestAirbaseSteps(t *testing.T) {
	assert.Nil(t, InitGameWithLogLevel(0, logrus.WarnLevel))
	ab := NewAirbase("Airbase 1", "usa", hexagon.NewHexagon(15, 15))
	Globals.World.AddAirbase(ab.AirbaseId)

	ac1 := NewAircraft("F14", "Default", "usa")
	ac1.FillSeatsWithNewPilots(OF1)
	ac1.Damage = append(ac1.Damage, []DamageType{DamageTypeFuselage, DamageTypeCockpit}...)
	err := ac1.FSM.Event(AcEventRepair)
	assert.Nil(t, err)

	ac1.AssignToAB(ab.AirbaseId)
	Globals.World.Step(10)
	Globals.World.Step(1000)

	fmt.Println(Globals.EventList.String())
}

func TestAirbaseAndAircraftsInTheAir(t *testing.T) {
	assert.Nil(t, InitGameWithLogLevel(0, logrus.WarnLevel))
	ab := NewAirbase("Airbase 1", "usa", hexagon.NewHexagon(15, 15))
	ac1 := NewAircraft("F14", "Default", "usa")
	ac1.FillSeatsWithNewPilots(OF1)
	ac1.AssignToAB(ab.AirbaseId)
	ab.AddToParkingArea(ac1.AircraftId)
	Globals.World.Step(1)
	assert.Equal(t, 1, len(ab.ParkingArea))
	assert.Equal(t, 0, len(Globals.AircraftsInTheAir))
	ac1.Waypoints = []hexagon.HexPosition{{
		Column: 15,
		Row:    15,
	}}
	Globals.World.Step(1)
	assert.Equal(t, 0, len(ab.ParkingArea))
	assert.Equal(t, 1, len(Globals.AircraftsInTheAir))
}

// TestStartAndLand starts and lands an aircraft.
func TestStartAndLand(t *testing.T) {
	assert.Nil(t, InitGameWithLogLevel(0, logrus.WarnLevel))
	ab := NewAirbase("Airbase 1", "usa", hexagon.NewHexagon(15, 15))
	ac1 := NewAircraft("F14", "Default", "usa")
	ac1.FillSeatsWithNewPilots(OF1)
	ac1.AssignToAB(ab.AirbaseId)
	ab.AddToParkingArea(ac1.AircraftId)
	waypoints := []hexagon.HexPosition{{17, 17}, {15, 15}}
	ac1.SetWaypoints(waypoints)
	assert.Equal(t, 4, len(ac1.CalculatedPath))
	assert.Equal(t, 1, len(ab.ParkingArea))
	assert.Equal(t, 0,
		len(Globals.AircraftsInTheAir))
	assert.Equal(t, hexagon.HexPosition{
		Column: 15,
		Row:    15,
	}, ac1.CurrentPosition)
	Globals.World.Step(1)
	assert.Equal(t, 0, len(ab.ParkingArea))
	assert.Equal(t, 1, len(Globals.AircraftsInTheAir))
	Globals.World.Step(1)
	assert.Equal(t, hexagon.HexPosition{
		Column: 15,
		Row:    16,
	}, ac1.CurrentPosition)
	Globals.World.Step(1)
	assert.Equal(t, hexagon.HexPosition{
		Column: 16,
		Row:    16,
	}, ac1.CurrentPosition)
	Globals.World.Step(1)
	assert.Equal(t, hexagon.HexPosition{
		Column: 17,
		Row:    17,
	}, ac1.CurrentPosition)
	Globals.World.Step(1)
	assert.Equal(t, hexagon.HexPosition{
		Column: 15,
		Row:    15,
	}, ac1.Destination)
	Globals.World.Step(1)
	assert.Equal(t, hexagon.HexPosition{
		Column: 16,
		Row:    16,
	}, ac1.CurrentPosition)
	Globals.World.Step(1)
	assert.Equal(t, hexagon.HexPosition{
		Column: 15,
		Row:    16,
	}, ac1.CurrentPosition)
	Globals.World.Step(1)
	assert.Equal(t, hexagon.HexPosition{
		Column: 15,
		Row:    15,
	}, ac1.CurrentPosition)
	Globals.World.Step(1)
	assert.Equal(t, hexagon.HexPosition{
		Column: 15,
		Row:    15,
	}, ac1.CurrentPosition)
	assert.Equal(t, 1, len(ab.ParkingArea))
	//assert.Equal(t, 0, len(Globals.AircraftsInTheAir))
}
