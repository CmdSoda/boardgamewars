package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/countrycodes"
	"github.com/CmdSoda/boardgamewars/internal/nato"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Nil(t, InitGame())
	wp := NewWarParty("USA", countrycodes.USA, Blue)
	ac := NewAircraft("F14", "Default", &wp.WarPartyId)
	pl := NewPilots(ac.GetParameters().Seats, wp.WarPartyId, nato.OF1)
	ac.FillSeatsWith(pl)
	wp.Aircrafts[ac.AircraftId] = ac
	fmt.Println(wp)
}
