package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/countrycodes"
	"testing"
)

func TestAirbase(t *testing.T) {
	InitGame()
	ab := NewAirbase("Fallujah AB", countrycodes.USA, Position{3, 5})
	ab.CreateAircrafts("F14", "Default", countrycodes.USA, 6)
	fmt.Println(ab.String())
}
