package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/hexagon"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewWorld(t *testing.T) {
	w := NewWorld()
	fmt.Println(w)
}

func TestStep(t *testing.T) {
	assert.Nil(t, InitGameWithLogLevel(0, logrus.WarnLevel))
	Step(3)
	assert.Equal(t, StepTime(3), Globals.World.CurrentStep)
	Step(7)
	assert.Equal(t, StepTime(10), Globals.World.CurrentStep)
}

func TestAirbaseSteps(t *testing.T) {
	assert.Nil(t, InitGameWithLogLevel(0, logrus.WarnLevel))
	ab := NewAirbase("Airbase 1", "usa", hexagon.NewHexagon(15, 15))
	Globals.World.AddAirbase(ab.AirbaseId)
	Globals.World.Step(1)
}
