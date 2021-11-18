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
	for i := 0; i < 2; i++ {
		b := game.NewAircraft("F14", "Default", game.WarPartyIdUSA)
		bpl := game.NewPilots(2, game.WarPartyIdUSA, nato.OF1)
		b.FillSeatsWith(bpl)
		ds.AddBlue(b.AircraftId)
		r := game.NewAircraft("F14", "Default", game.WarPartyIdRussia)
		rpl := game.NewPilots(2, game.WarPartyIdRussia, nato.OF1)
		r.FillSeatsWith(rpl)
		ds.AddRed(r.AircraftId)
	}
	d := game.NewDogfight(ds)
	d.DistributeAircraftsToGroups()
	d.Simulate()
	fmt.Println(d)
}
