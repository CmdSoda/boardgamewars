package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/nato"
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

func TestAircraftStep(t *testing.T) {
	assert.Nil(t, InitGameWithLogLevel(0, logrus.WarnLevel))
	ac1 := NewAircraft("F14", "Default", WarPartyIdUSA)
	ac1.FillSeatsWithNewPilots(nato.OF1)
	Globals.World.Add(ac1)
	Globals.World.Step(3)
	assert.Equal(t, StepTime(3), ac1.StepsTaken)
}
