package military

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/randomizer"
	"testing"
)

func TestRanks(t *testing.T) {
	randomizer.Init()
	frl := NewRank(AirForceRAF, NatoOfficerCodeOF1)
	fmt.Println(frl)
}
