package main

import (
	"fmt"
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
		acblue := game.NewAircraft(parameters.Name, "Default", game.WarPartyIdUSA)
		acblue.FillSeatsWithNewPilots(nato.OF1)
		fmt.Println(acblue)
		for _, parameters2 := range game.Globals.AllAircraftParameters {
			if parameters.Name != parameters2.Name {
				acred := game.NewAircraft(parameters2.Name, "Default", game.WarPartyIdRussia)
				acred.FillSeatsWithNewPilots(nato.OF1)
				for i := 0; i < 10000; i++ {
					// Blau reparieren und bewaffnen
					ds := game.NewDogfightSetup()
					ds.AddBlue(acblue.AircraftId)
					acblue.ReviveAndRepair()
					acblue.Rearm()
					// Rot erstellen
					ds.AddRed(acred.AircraftId)
					// Simulation
					d := game.NewDogfight(ds)
					d.DistributeAircraftsToGroups()
					//fmt.Println(d)
					d.Simulate()
				}
			}
		}
	}
	game.Globals.Statistics.DmgVsA.Dump()
	game.Globals.Statistics.W2A2C.Dump()
	game.Globals.Statistics.WVsA.Dump()
}
