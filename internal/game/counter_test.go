package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/nato"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCounter(t *testing.T) {
	c := NewCounter(CounterTypeAircraft, nil)
	fmt.Println(c)
}

func TestCounterType_String(t *testing.T) {
	c := NewCounter(99, nil)
	assert.Equal(t, "Unknown", c.Type.String())
}

func Test2(t *testing.T) {
	assert.Nil(t, InitGame(0))
	ac := NewAircraft("F14", "Default", WarPartyIdUSA)
	c := NewCounter(CounterTypeAircraft, ac)
	assert.Equal(t, "Aircraft", c.Type.String())
	ac2 := c.Object.(*Aircraft)
	assert.Equal(t, ac.AircraftId, ac2.AircraftId)
}

func TestCounterAircraft(t *testing.T) {
	assert.Nil(t, InitGameWithLogLevel(0, logrus.WarnLevel))
	ac := NewAircraft("F14", "Default", WarPartyIdUSA)
	ac.FillSeatsWithNewPilots(nato.OF1)
	c := NewCounter(CounterTypeAircraft, ac)
	fmt.Println(c)
}
