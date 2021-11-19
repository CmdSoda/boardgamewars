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
	Globals.Statistics.WVsA.Dump()
}

func TestWin1(t *testing.T) {
	assert.Nil(t, InitGameWithLogLevel(0, logrus.WarnLevel))
	s := NewStatistics()
	ac1 := NewAircraft("F14", "Default", WarPartyIdUSA)
	ac1.FillSeatsWithNewPilots(nato.OF1)
	ac2 := NewAircraft("MiG-29", "Default", WarPartyIdRussia)
	ac2.FillSeatsWithNewPilots(nato.OF1)
	ac3 := NewAircraft("F5", "Default", WarPartyIdRussia)
	ac3.FillSeatsWithNewPilots(nato.OF1)
	ac4 := NewAircraft("F5", "Default", WarPartyIdRussia)
	ac4.FillSeatsWithNewPilots(nato.OF1)
	ac5 := NewAircraft("F14", "Default", WarPartyIdRussia)
	ac5.FillSeatsWithNewPilots(nato.OF1)
	s.WVsA.Win(ac1.AircraftId, ac2.AircraftId, WinTypeWon)
	s.WVsA.Win(ac1.AircraftId, ac2.AircraftId, WinTypeWon)
	s.WVsA.Win(ac2.AircraftId, ac1.AircraftId, WinTypeWon)
	s.WVsA.Win(ac2.AircraftId, ac3.AircraftId, WinTypeWon)
	s.WVsA.Win(ac2.AircraftId, ac3.AircraftId, WinTypeDraw)
	s.WVsA.Win(ac3.AircraftId, ac4.AircraftId, WinTypeDraw)
	s.WVsA.Win(ac4.AircraftId, ac3.AircraftId, WinTypeDraw)
	s.WVsA.Win(ac4.AircraftId, ac5.AircraftId, WinTypeDraw)
	s.WVsA.Win(ac5.AircraftId, ac4.AircraftId, WinTypeDraw)
	assert.Equal(t, 4, len(s.WVsA))
	assert.Equal(t, ac1.AircraftParametersId, s.WVsA[0].AC1Params.AircraftParametersId)
	assert.Equal(t, ac2.AircraftParametersId, s.WVsA[0].AC2Params.AircraftParametersId)

	assert.Equal(t, 2, s.WVsA[0].AC1Won)
	assert.Equal(t, 1, s.WVsA[0].AC2Won)
	assert.Equal(t, 0, s.WVsA[0].Draw)

	assert.Equal(t, 1, s.WVsA[1].AC1Won)
	assert.Equal(t, 1, s.WVsA[1].Draw)

	assert.Equal(t, 0, s.WVsA[2].AC1Won)
	assert.Equal(t, 0, s.WVsA[2].AC2Won)
	assert.Equal(t, 2, s.WVsA[2].Draw)

	s.WVsA.Dump()
}

func TestDamageMapAdd(t *testing.T) {
	assert.Nil(t, InitGameWithLogLevel(0, logrus.WarnLevel))
	ac1 := NewAircraft("F14", "Default", WarPartyIdUSA)
	ac1.FillSeatsWithNewPilots(nato.OF1)
	ac2 := NewAircraft("MiG-29", "Default", WarPartyIdRussia)
	ac2.FillSeatsWithNewPilots(nato.OF1)
	for i := 0; i < 1000; i++ {
		ac1.Rearm()
		wps, _ := ac1.GetBestDogfightingWeapon()
		ac2.ReviveAndRepair()
		ac2.DoDamageWith(wps)
	}
	Globals.Statistics.DmgVsA.Dump()
}
