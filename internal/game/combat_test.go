package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/countrycodes"
	"github.com/CmdSoda/boardgamewars/internal/nato"
	"testing"
)

func TestDogfight(t *testing.T) {
	InitGame()

	fighter1 := NewAircraftManned("F14", "Default", countrycodes.UK, nato.OF2)
	fighter2 := NewAircraftManned("F14", "Default", countrycodes.UK, nato.OF1)

	var dr1 DogfightResult
	var dr2 DogfightResult
	dfp1 := &DogfightParameters{
		Aircraft:             *fighter1,
		DogfightResult:       dr1,
		LastDogfightPosition: 0,
	}
	dfp2 := &DogfightParameters{
		Aircraft:             *fighter2,
		DogfightResult:       dr2,
		LastDogfightPosition: 0,
	}
	Dogfight(dfp1, dfp2)
	fmt.Println(dfp1.DogfightResult)
	fmt.Println(dfp2.DogfightResult)
}

func TestMoreRounds(t *testing.T) {
	InitGame()

	fighter1 := NewAircraftManned("F14", "Default", countrycodes.UK, nato.OF2)
	fighter2 := NewAircraftManned("F14", "Default", countrycodes.UK, nato.OF1)

	drl1, drl2 := Sim10Rounds(fighter1, fighter2)
	fmt.Println(drl1)
	fmt.Println(drl2)
}
