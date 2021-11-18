package game

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStatistics_Hit(t *testing.T) {
	assert.Nil(t, InitGameWithLogLevel(0, logrus.WarnLevel))
	s := NewStatistics()
	ac := NewAircraft("F14", "Default", WarPartyIdUSA)
	s.Hit("Aim-7", ac.AircraftId)
	s.Hit("Aim-7", ac.AircraftId)
	s.Hit("Aim-7", ac.AircraftId)
	s.NotHit("Aim-7", ac.AircraftId)
	s.Hit("Aim-9", ac.AircraftId)
	assert.Equal(t, 3, s.W2A2C["Aim-7"][ac.AircraftParametersId]["Default"].Hit)
	assert.Equal(t, 1, s.W2A2C["Aim-7"][ac.AircraftParametersId]["Default"].NotHit)
	assert.Equal(t, 1, s.W2A2C["Aim-9"][ac.AircraftParametersId]["Default"].Hit)
	assert.Equal(t, 0, s.W2A2C["Aim-9"][ac.AircraftParametersId]["Default"].NotHit)

	s.W2A2C.Dump()
}
