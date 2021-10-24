package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/countrycodes"
	"github.com/CmdSoda/boardgamewars/internal/military"
	"github.com/CmdSoda/boardgamewars/internal/randomizer"
	"testing"
)

func TestNewPilot(t *testing.T) {
	randomizer.Init()
	p := NewPilot(countrycodes.UK, military.NatoOfficerCodeOF2)
	fmt.Println(p)
}
