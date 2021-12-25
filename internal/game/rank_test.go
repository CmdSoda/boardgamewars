package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/randomizer"
	"testing"
)

func TestRanks(t *testing.T) {
	randomizer.Init(0)
	frl := NewRank("uk", OF1)
	fmt.Println(frl)
}
