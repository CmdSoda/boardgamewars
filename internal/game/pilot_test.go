package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/countrycodes"
	"github.com/CmdSoda/boardgamewars/internal/nato"
	"testing"
)

func TestNewPilot(t *testing.T) {
	InitGame()
	p := NewPilot(countrycodes.UK, nato.OF2)
	fmt.Println(p)
}
