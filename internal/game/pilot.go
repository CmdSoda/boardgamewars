package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/countrycodes"
	"github.com/CmdSoda/boardgamewars/internal/military"
	"github.com/CmdSoda/boardgamewars/internal/namegenerator"
)

type Pilot struct {
	Name    string
	Country string
	military.FlightRank
}

func (p Pilot) String() string {
	return fmt.Sprintf("%s (%s) [%s]", p.Name, p.FlightRank, p.Country)
}

func NewPilot(cc countrycodes.Code, ofc military.NatoOfficerCode) Pilot {
	return Pilot{
		Name:       namegenerator.CreateMaleFullName(cc),
		Country:    cc.String(),
		FlightRank: military.NewRank(military.AirForceRAF, ofc),
	}
}
