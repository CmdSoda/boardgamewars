package military

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/countrycodes"
	"github.com/CmdSoda/boardgamewars/internal/nato"
	"github.com/CmdSoda/boardgamewars/internal/randomizer"
)

type FlightRank struct {
	Name  string
	Short string
	nato.Code
}

func (f FlightRank) String() string {
	if f.Short != "" {
		return fmt.Sprintf("%s", f.Short)
	} else {
		return f.Name
	}
}

type FlightRankList []FlightRank

func (frl FlightRankList) RandomPick() FlightRank {
	roll := randomizer.Roll1DN(len(frl))
	return frl[roll-1]
}

func GetAllRanksWithNatoCode(list []FlightRank, noc nato.Code) FlightRankList {
	frl := make(FlightRankList, 0)
	for _, rank := range list {
		if rank.Code == noc {
			frl = append(frl, rank)
		}
	}
	return frl
}

func GetFlightRank(list []FlightRank, noc nato.Code) FlightRank {
	frl := GetAllRanksWithNatoCode(list, noc)
	return frl.RandomPick()
}

var TheFlightRanks = map[countrycodes.Code][]FlightRank {
	countrycodes.UK: PilotRanksUK,
	countrycodes.Germany: PilotRanksGermany,
	countrycodes.USA: PilotRanksUSA,
	countrycodes.Russia: PilotRanksUSA,
}

func NewRank(cc countrycodes.Code, noc nato.Code) FlightRank {
	return GetFlightRank(TheFlightRanks[cc], noc)
}
