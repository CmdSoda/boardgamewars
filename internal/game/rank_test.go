package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/nato"
	"github.com/CmdSoda/boardgamewars/internal/randomizer"
	"testing"
)

func TestRanks(t *testing.T) {
	randomizer.Init(0)
	frl := NewRank("uk", nato.OF1)
	fmt.Println(frl)
}
