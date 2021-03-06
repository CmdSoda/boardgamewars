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
	s.WeaponPerformance.Hit("Aim-7", ac.AircraftId)
	s.WeaponPerformance.Hit("Aim-7", ac.AircraftId)
	s.WeaponPerformance.Hit("Aim-7", ac.AircraftId)
	s.WeaponPerformance.NotHit("Aim-7", ac.AircraftId)
	s.WeaponPerformance.Hit("Aim-9", ac.AircraftId)
	assert.Equal(t, 3, s.WeaponPerformance["Aim-7"][ac.AircraftParametersId]["Default"].Hit)
	assert.Equal(t, 1, s.WeaponPerformance["Aim-7"][ac.AircraftParametersId]["Default"].NotHit)
	assert.Equal(t, 1, s.WeaponPerformance["Aim-9"][ac.AircraftParametersId]["Default"].Hit)
	assert.Equal(t, 0, s.WeaponPerformance["Aim-9"][ac.AircraftParametersId]["Default"].NotHit)
	s.WeaponPerformance.Dump()
}

func TestStatistics2(t *testing.T) {
	assert.Nil(t, InitGameWithLogLevel(0, logrus.WarnLevel))
	ds := NewDogfightSetup()
	for i := 0; i < 500; i++ {
		b := NewAircraft("F14", "Default", WarPartyIdUSA)
		b.FillSeatsWithNewPilots(OF1)
		ds.AddBlue(b.AircraftId)
		r := NewAircraft("MiG-29", "Default", WarPartyIdRussia)
		r.FillSeatsWithNewPilots(OF1)
		ds.AddRed(r.AircraftId)
	}
	d := NewDogfight(ds)
	assert.True(t, d.DistributeAircraftsToGroups())
	for round := 0; round < 40; round++ {
		d.Simulate()
	}
	Globals.Statistics.WeaponPerformance.Dump()
	Globals.Statistics.AircraftVsAircraft.Dump()
}

func TestWin1(t *testing.T) {
	assert.Nil(t, InitGameWithLogLevel(0, logrus.WarnLevel))
	s := NewStatistics()
	ac1 := NewAircraft("F14", "Default", WarPartyIdUSA)
	ac1.FillSeatsWithNewPilots(OF1)
	ac2 := NewAircraft("MiG-29", "Default", WarPartyIdRussia)
	ac2.FillSeatsWithNewPilots(OF1)
	ac3 := NewAircraft("F5", "Default", WarPartyIdRussia)
	ac3.FillSeatsWithNewPilots(OF1)
	ac4 := NewAircraft("F5", "Default", WarPartyIdRussia)
	ac4.FillSeatsWithNewPilots(OF1)
	ac5 := NewAircraft("F14", "Default", WarPartyIdRussia)
	ac5.FillSeatsWithNewPilots(OF1)
	s.AircraftVsAircraft.Win(ac1.AircraftId, ac2.AircraftId, WinTypeWon)
	s.AircraftVsAircraft.Win(ac1.AircraftId, ac2.AircraftId, WinTypeWon)
	s.AircraftVsAircraft.Win(ac2.AircraftId, ac1.AircraftId, WinTypeWon)
	s.AircraftVsAircraft.Win(ac2.AircraftId, ac3.AircraftId, WinTypeWon)
	s.AircraftVsAircraft.Win(ac2.AircraftId, ac3.AircraftId, WinTypeDraw)
	s.AircraftVsAircraft.Win(ac3.AircraftId, ac4.AircraftId, WinTypeDraw)
	s.AircraftVsAircraft.Win(ac4.AircraftId, ac3.AircraftId, WinTypeDraw)
	s.AircraftVsAircraft.Win(ac4.AircraftId, ac5.AircraftId, WinTypeDraw)
	s.AircraftVsAircraft.Win(ac5.AircraftId, ac4.AircraftId, WinTypeDraw)
	assert.Equal(t, 4, len(s.AircraftVsAircraft))
	assert.Equal(t, ac1.AircraftParametersId, s.AircraftVsAircraft[0].AC1Params.AircraftParametersId)
	assert.Equal(t, ac2.AircraftParametersId, s.AircraftVsAircraft[0].AC2Params.AircraftParametersId)

	assert.Equal(t, 2, s.AircraftVsAircraft[0].AC1Stats.Won)
	assert.Equal(t, 1, s.AircraftVsAircraft[0].AC2Stats.Won)
	assert.Equal(t, 0, s.AircraftVsAircraft[0].AC1Stats.Draw)

	assert.Equal(t, 1, s.AircraftVsAircraft[1].AC1Stats.Won)
	assert.Equal(t, 1, s.AircraftVsAircraft[1].AC1Stats.Draw)

	assert.Equal(t, 0, s.AircraftVsAircraft[2].AC1Stats.Won)
	assert.Equal(t, 0, s.AircraftVsAircraft[2].AC2Stats.Won)
	assert.Equal(t, 2, s.AircraftVsAircraft[2].AC1Stats.Draw)

	s.AircraftVsAircraft.Dump()
}

func TestDamageMapAdd(t *testing.T) {
	assert.Nil(t, InitGameWithLogLevel(0, logrus.WarnLevel))
	ac1 := NewAircraft("F14", "Default", WarPartyIdUSA)
	ac1.FillSeatsWithNewPilots(OF1)
	ac2 := NewAircraft("MiG-29", "Default", WarPartyIdRussia)
	ac2.FillSeatsWithNewPilots(OF1)
	for i := 0; i < 1000; i++ {
		ac1.Rearm()
		wps, _ := ac1.GetBestDogfightingWeapon()
		ac2.ReviveAndRepair()
		ac2.DoDamageWith(wps)
	}
	Globals.Statistics.WeaponPerformance.Dump()
}

func TestFreeForAll(t *testing.T) {
	assert.Nil(t, InitGameWithLogLevel(0, logrus.WarnLevel))
	FreeForAll()
	Globals.Statistics.AircraftVsAircraft.Dump()
	Globals.Statistics.WeaponPerformance.Dump()
}
