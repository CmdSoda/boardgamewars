package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/countrycodes"
	"github.com/CmdSoda/boardgamewars/internal/nato"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAircraftPilots(t *testing.T) {
	assert.Nil(t, InitGame())
	ac := NewAircraftManned("F14", "Default", countrycodes.UK, nato.OF2)
	fmt.Println(ac)
	ac = NewAircraftManned("F14", "Default", countrycodes.Germany, nato.OF2)
	fmt.Println(ac)
	ac = NewAircraftManned("F14", "Default", countrycodes.USA, nato.OF2)
	fmt.Println(ac)
	ac = NewAircraftManned("F14", "Default", countrycodes.Russia, nato.OF2)
	fmt.Println(ac)
	fmt.Println(Globals.PilotRoster)
}

func TestAssign(t *testing.T) {
	assert.Nil(t, InitGame())
	wrongid := uuid.New()
	ac := NewAircraft("F14", "Default", countrycodes.UK)
	assert.Equal(t, false, ac.AssignToAB(wrongid))
	nellis := NewAirbase("Nellis AB", countrycodes.USA, Position{6, 9})
	assert.Equal(t, true, ac.AssignToAB(nellis.Id))
}
