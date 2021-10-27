package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/countrycodes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAirbase(t *testing.T) {
	assert.Nil(t, InitGame())
	ab := NewAirbase("Fallujah AB", countrycodes.USA, Position{3, 5})
	ab.CreateAircrafts("F14", "Default", countrycodes.USA, 6)
	fmt.Println(ab.String())
	fmt.Println(Globals.AirbaseList)
}

func TestHangar(t *testing.T) {
	assert.Nil(t, InitGame())
	ab := NewAirbase("Fallujah AB", countrycodes.USA, Position{3, 5})
	ac := NewAircraft("F14", "Default", countrycodes.USA)
	ab.AddToHangar(ac)
	fmt.Println(ab)
}
