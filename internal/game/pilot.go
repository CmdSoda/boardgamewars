package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/countrycodes"
	"github.com/CmdSoda/boardgamewars/internal/military"
	"github.com/CmdSoda/boardgamewars/internal/namegenerator"
	"github.com/CmdSoda/boardgamewars/internal/randomizer"
)

type Gender int

const (
	GenderMale   Gender = 0
	GenderFemale Gender = 1
)

func (g Gender) String() string {
	if g == GenderMale {
		return "Male"
	} else {
		return "Female"
	}
}

type Pilot struct {
	Name string
	Gender
	countrycodes.Code
	military.FlightRank
}

func (p Pilot) String() string {
	return fmt.Sprintf("%s(%s) (%s) [%s]", p.Name, p.Gender, p.FlightRank, p.Code.String())
}

func NewPilot(cc countrycodes.Code, ofc military.NatoOfficerCode) Pilot {
	var af military.AirForce
	switch cc {
	case countrycodes.UK:
		af = military.AirForceRAF
	case countrycodes.Germany:
		af = military.AirForceGAF
	case countrycodes.USA:
		af = military.AirForceUSAF
	}

	var g Gender
	gr := randomizer.Roll1D10()
	if gr >= 2 {
		g = GenderMale
	} else {
		g = GenderFemale
	}

	var fn func(countrycodes.Code) string
	if g == GenderMale {
		fn = namegenerator.CreateMaleFullName
	} else {
		fn = namegenerator.CreateFemaleFullName
	}

	return Pilot{
		Name:       fn(cc),
		Gender:     g,
		Code:       cc,
		FlightRank: military.NewRank(af, ofc),
	}
}
