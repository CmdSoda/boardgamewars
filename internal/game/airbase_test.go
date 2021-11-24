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
	ab.AddToHangar(ac.AircraftId)
	fmt.Println(ab)
}

func TestNewAirbase(t *testing.T) {
	assert.Nil(t, InitGameWithLogLevel(0, logrus.WarnLevel))
	ab := NewAirbase("Airbase 1", WarPartyIdUSA, hexagon.NewHexagon(15, 15))
	assert.Equal(t, 1, len(Globals.AllAirbases))
	_, exist := Globals.AllAirbases[ab.AirbaseId]
	assert.True(t, exist)
}
