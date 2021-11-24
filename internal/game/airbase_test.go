package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/hexagon"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAirbase(t *testing.T) {
	assert.Nil(t, InitGame(0))
	ab := NewAirbase("Fallujah AB", WarPartyIdUSA, hexagon.HexPosition{Column: 3, Row: 5})
	ab.CreateAircrafts("F14", "Default", WarPartyIdUSA, 6)
	fmt.Println(ab.String())
	fmt.Println(Globals.AllAirbases)
}

func TestHangar(t *testing.T) {
	assert.Nil(t, InitGame(0))
	ab := NewAirbase("Fallujah AB", WarPartyIdUSA, hexagon.HexPosition{Column: 3, Row: 5})
	ac := NewAircraft("F14", "Default", WarPartyIdUSA)
	ab.AddToParkingArea(ac.AircraftId)
	fmt.Println(ab)
}

func TestNewAirbase(t *testing.T) {
	assert.Nil(t, InitGameWithLogLevel(0, logrus.WarnLevel))
	ab := NewAirbase("Airbase 1", WarPartyIdUSA, hexagon.NewHexagon(15, 15))
	assert.Equal(t, 1, len(Globals.AllAirbases))
	_, exist := Globals.AllAirbases[ab.AirbaseId]
	assert.True(t, exist)
}

func TestMoveAircraftsToMaintenance(t *testing.T) {
	assert.Nil(t, InitGameWithLogLevel(0, logrus.WarnLevel))
	ab := NewAirbase("Airbase", WarPartyIdUSA, hexagon.NewHexagon(15, 15))
	ab.MaxMaintainenanceSlots = 4
	ac1 := NewAircraft("F14", "Default", WarPartyIdUSA)
	ac2 := NewAircraft("F14", "Default", WarPartyIdUSA)
	ac3 := NewAircraft("F14", "Default", WarPartyIdUSA)
	ac4 := NewAircraft("F14", "Default", WarPartyIdUSA)
	ac5 := NewAircraft("F14", "Default", WarPartyIdUSA)
	ac6 := NewAircraft("F14", "Default", WarPartyIdUSA)
	ab.AddToParkingArea(ac1.AircraftId)
	ab.AddToParkingArea(ac2.AircraftId)
	ab.AddToParkingArea(ac3.AircraftId)
	ab.AddToParkingArea(ac4.AircraftId)
	ab.AddToParkingArea(ac5.AircraftId)
	ab.AddToParkingArea(ac6.AircraftId)
	ac2.AddDamage([]DamageType{DamageTypeFuselage, DamageTypeRudder})
	ac3.AddDamage([]DamageType{DamageTypeFuselage, DamageTypeRudder})
	ac4.AddDamage([]DamageType{DamageTypeFuselage, DamageTypeRudder})
	ac5.AddDamage([]DamageType{DamageTypeFuselage, DamageTypeRudder})
	ac6.AddDamage([]DamageType{DamageTypeFuselage, DamageTypeRudder})
	ab.Step(1)
	assert.Equal(t, 2, len(ab.ParkingArea))
	assert.Equal(t, 4, len(ab.MaintenanceArea))
	assert.Equal(t, ac1.AircraftId, ab.ParkingArea[0])
	assert.Equal(t, ac6.AircraftId, ab.ParkingArea[1])
	assert.Equal(t, ac2.AircraftId, ab.MaintenanceArea[0])
	assert.Equal(t, ac3.AircraftId, ab.MaintenanceArea[1])
	assert.Equal(t, ac4.AircraftId, ab.MaintenanceArea[2])
	assert.Equal(t, ac5.AircraftId, ab.MaintenanceArea[3])
}
