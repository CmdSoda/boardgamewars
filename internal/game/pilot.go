package game

import (
	"fmt"
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

type PilotsMap map[PilotId]*Pilot

type Pilot struct {
	PilotId
	Name        string
	CountryName // Gehört diesem Land an
	Gender
	Background PilotBackground
	FlightRank
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
		p.Background.CountryName,
		p.PilotStats.Sorties,
		p.PilotStats.Hits,
		p.PilotStats.Kills,
		p.Background.Age,
		p.Background.Born,
		p.Background.HomeAirBase)
}

func (p Pilot) Short() string {
	return p.Name
}

// RollAge
// https://www.operationmilitarykids.org/air-force-age-limits/
func RollAge(ofc Code) int {
	switch ofc {
	case OF1:
		return randomizer.Roll(22, 24)
	case OF2:
		return randomizer.Roll(25, 28)
	case OF3:
		return randomizer.Roll(27, 31)
	case OF4:
		return randomizer.Roll(29, 35)
	case OF5:
		return randomizer.Roll(32, 39)
	}
	return 0
}

func NewPilot(country CountryName, ofc Code) *Pilot {
	var g Gender

	wp, exist := Globals.CountryDataMap[country]
	if exist == false {
		panic("country does not exist")
	}

	gr := randomizer.Roll1D10()
	if gr >= 2 {
		g = GenderMale
	} else {
		g = GenderFemale
	}

	ng := Generator{}
	var cd *CountryDataItem
	cd, exist = Globals.CountryDataMap[wp.CountryName]
	if !exist {
		Log.Panicf("data for %s not found", wp.CountryName)
		return nil
	}
	ng.AddNameSet(&cd.NameSet)

	np := Pilot{
		Name:        ng.CreateFullName(g == GenderMale, wp.CountryName),
		CountryName: country,
		PilotId:     PilotId(uuid.New()),
		Gender:      g,
		Background: PilotBackground{
			CountryName: wp.CountryName,
			Born:        ng.CreateCityName(wp.CountryName),
			Age:         RollAge(ofc),
			HomeAirBase: ng.CreateAirForceBaseName(wp.CountryName),
		},
		FlightRank: NewRank(wp.CountryName, ofc),
	}

	Globals.CountryDataMap[country].Pilots = append(Globals.CountryDataMap[country].Pilots, np.PilotId)
	Globals.AllPilots[np.PilotId] = &np
	Log.Infof("new pilot created: %s", np.Short())
	return &np
}

func NewPilots(count int, country CountryName, ofc Code) []PilotId {
	var pilots []PilotId
	for i := 0; i < count; i++ {
		np := NewPilot(country, ofc)
		pilots = append(pilots, np.PilotId)
	}
	return pilots
}
