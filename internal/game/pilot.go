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
	Reflexes  int // Reflexe im Dogfight: -2, -1, 0, +1, +2
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

func RollAge(ofc nato.Code) int {
	return 18 + randomizer.Roll1DN(5) + 4*(int(ofc)-1)
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
