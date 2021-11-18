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
				ds := game.NewDogfightSetup()

				// Blau reparieren und bewaffnen
				acblue.ReviveAndRepair()
				acblue.Rearm()
				ds.AddBlue(acblue.AircraftId)

				// Rot erstellen
				acred := game.NewAircraft(parameters2.Name, "Default", game.WarPartyIdRussia)
				acred.FillSeatsWithNewPilots(nato.OF1)
				fmt.Println(acred)
				ds.AddRed(acred.AircraftId)

				// Simulation
				d := game.NewDogfight(ds)
				d.DistributeAircraftsToGroups()
				fmt.Println(d)
				d.Simulate()
			}
		}
	}
}
