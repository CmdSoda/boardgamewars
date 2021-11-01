package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/countrycodes"
	"github.com/CmdSoda/boardgamewars/internal/nato"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPilot(t *testing.T) {
	assert.Nil(t, InitGame())
	wp := NewWarParty("USA", countrycodes.Germany, Blue)
	p := NewPilot(wp.WarPartyId, nato.OF5)
	fmt.Println(p)
}

func TestNewPilots(t *testing.T) {
	assert.Nil(t, InitGame())
	wp := NewWarParty("USA", countrycodes.Russia, Blue)
	pl := NewPilots(3, wp.WarPartyId, nato.OF1)
	fmt.Println(pl)
}
