package main

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/game"
	"github.com/CmdSoda/boardgamewars/internal/nato"
	"github.com/sirupsen/logrus"
)

func main() {
	err := game.InitGameWithLogLevel(0, logrus.ErrorLevel)
	if err != nil {
		panic("Could not init game: " + err.Error())
		return
	}
	ds := game.NewDogfightSetup()
	b := game.NewAircraft("F14", "Default", game.WarPartyIdUSA)
	b.FillSeatsWithNewPilots(nato.OF1)
	fmt.Println(b)
	ds.AddBlue(b.AircraftId)
	r := game.NewAircraft("MiG-29", "Default", game.WarPartyIdRussia)
	r.FillSeatsWithNewPilots(nato.OF1)
	fmt.Println(r)
	ds.AddRed(r.AircraftId)
	d := game.NewDogfight(ds)
	d.DistributeAircraftsToGroups()
	d.Simulate()
	fmt.Println(d)
}
