package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/hexagon"
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
