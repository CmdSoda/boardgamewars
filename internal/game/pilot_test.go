package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/countrycodes"
	"github.com/CmdSoda/boardgamewars/internal/military"
	"testing"
)

func TestNewPilot(t *testing.T) {
	InitGame()
	p := NewPilot(countrycodes.UK, military.NatoOfficerCodeOF2)
	fmt.Println(p)
}
