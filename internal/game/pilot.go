package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/countrycodes"
	"github.com/CmdSoda/boardgamewars/internal/military"
	"github.com/CmdSoda/boardgamewars/internal/namegenerator"
	"github.com/CmdSoda/boardgamewars/internal/nato"
	"github.com/CmdSoda/boardgamewars/internal/randomizer"
	"github.com/google/uuid"
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

type PilotId uuid.UUID

type Pilot struct {
	PilotId
	Name string
	Gender
	Background PilotBackground
	military.FlightRank
	PilotStats
}

type PilotStats struct {
	Sorties   int
	Hits      int
	Kills     int
	Kia       bool
	Mia       bool
	XP        int
	Reflexes  int // Reflexe im ExecuteDogfight: -2, -1, 0, +1, +2
	Endurance int // Wieviel Stress kann der Pilot verkraften.
}

func (p Pilot) String() string {
	return fmt.Sprintf("%s(%s) (%s) [%s] (%d-%d-%d) %dyo (Born: %s HomeBase: %s)",
		p.Name,
		p.Gender,
		p.FlightRank,
		p.Background.Country.String(),
		p.PilotStats.Sorties,
		p.PilotStats.Hits,
		p.PilotStats.Kills,
		p.Background.Age,
		p.Background.Born,
		p.Background.HomeAirBase)
}

// RollAge
// https://www.operationmilitarykids.org/air-force-age-limits/
func RollAge(ofc nato.Code) int {
	switch ofc {
	case nato.OF1:
		return randomizer.Roll(22, 24)
	case nato.OF2:
		return randomizer.Roll(25, 28)
	case nato.OF3:
		return randomizer.Roll(27, 31)
	case nato.OF4:
		return randomizer.Roll(29, 35)
	case nato.OF5:
		return randomizer.Roll(32, 39)
	}
	return 0
}

func NewPilot(cc countrycodes.Code, ofc nato.Code) Pilot {
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
		Name:    fn(cc),
		PilotId: PilotId(uuid.New()),
		Gender:  g,
		Background: PilotBackground{
			Country:     cc,
			Born:        namegenerator.CreateCityName(cc),
			Age:         RollAge(ofc),
			HomeAirBase: namegenerator.CreateAirForceBaseName(cc),
		},
		FlightRank: military.NewRank(cc, ofc),
	}
}
