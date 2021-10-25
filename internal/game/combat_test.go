package game

import (
	"github.com/CmdSoda/boardgamewars/internal/countrycodes"
	"github.com/CmdSoda/boardgamewars/internal/nato"
	"testing"
)

func TestDogfight(t *testing.T) {
	InitGame()

	fighter1 := NewAircraftManned("F14", "Default", countrycodes.UK, nato.OF2)
	fighter2 := NewAircraftManned("F14", "Default", countrycodes.UK, nato.OF1)

	Dogfight(fighter1, fighter2)
}
