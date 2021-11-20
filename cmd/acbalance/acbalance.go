package main

import (
	"github.com/CmdSoda/boardgamewars/internal/game"
	"github.com/CmdSoda/boardgamewars/internal/nato"
	"github.com/sirupsen/logrus"
)

func main() {
	err := game.InitGameWithLogLevel(0, logrus.WarnLevel)
	if err != nil {
		panic("Could not init game: " + err.Error())
		return
	}
	for _, parameters := range game.Globals.AllAircraftParameters {
		// Blau erstellen
		acblue := game.NewAircraft(parameters.Name, "Default", game.WarPartyIdUSA)
		acblue.FillSeatsWithNewPilots(nato.OF1)
		for _, parameters2 := range game.Globals.AllAircraftParameters {
			if parameters.Name != parameters2.Name {
				// Rot erstellen
				acred := game.NewAircraft(parameters2.Name, "Default", game.WarPartyIdRussia)
				acred.FillSeatsWithNewPilots(nato.OF1)
				for i := 0; i < 10000; i++ {
					acblue.ReviveAndRepair()
					acblue.Rearm()
					ds := game.NewDogfightSetup()
					ds.AddBlue(acblue.AircraftId)
					// Blau reparieren und bewaffnen
					ds.AddRed(acred.AircraftId)
					// Simulation
					d := game.NewDogfight(ds)
					d.DistributeAircraftsToGroups()
					d.Simulate()
				}
			}
		}
	}
	game.Globals.Statistics.AircraftVsAircraft.Dump()
	game.Globals.Statistics.WeaponPerformance.Dump()
	game.Globals.Statistics.WeaponDmgCount.Dump()
}
