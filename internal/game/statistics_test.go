package game

import (
	"github.com/CmdSoda/boardgamewars/internal/nato"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStatistics_Hit(t *testing.T) {
	assert.Nil(t, InitGameWithLogLevel(0, logrus.WarnLevel))
	s := NewStatistics()
	ac := NewAircraft("F14", "Default", WarPartyIdUSA)
	s.W2A2C.Hit("Aim-7", ac.AircraftId)
	s.W2A2C.Hit("Aim-7", ac.AircraftId)
	s.W2A2C.Hit("Aim-7", ac.AircraftId)
	s.W2A2C.NotHit("Aim-7", ac.AircraftId)
	s.W2A2C.Hit("Aim-9", ac.AircraftId)
	assert.Equal(t, 3, s.W2A2C["Aim-7"][ac.AircraftParametersId]["Default"].Hit)
	assert.Equal(t, 1, s.W2A2C["Aim-7"][ac.AircraftParametersId]["Default"].NotHit)
	assert.Equal(t, 1, s.W2A2C["Aim-9"][ac.AircraftParametersId]["Default"].Hit)
	assert.Equal(t, 0, s.W2A2C["Aim-9"][ac.AircraftParametersId]["Default"].NotHit)
	s.W2A2C.Dump()
}

func TestStatistics2(t *testing.T) {
	assert.Nil(t, InitGameWithLogLevel(0, logrus.WarnLevel))
	ds := NewDogfightSetup()
	for i := 0; i < 500; i++ {
		b := NewAircraft("F14", "Default", WarPartyIdUSA)
		b.FillSeatsWithNewPilots(nato.OF1)
		ds.AddBlue(b.AircraftId)
		r := NewAircraft("MiG-29", "Default", WarPartyIdRussia)
		r.FillSeatsWithNewPilots(nato.OF1)
		ds.AddRed(r.AircraftId)
	}
	d := NewDogfight(ds)
	assert.True(t, d.DistributeAircraftsToGroups())
	for round := 0; round < 40; round++ {
		d.Simulate()
	}
	Globals.Statistics.W2A2C.Dump()
}